package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/mavr/anonymous-mail/pkg/anonbot/anonbot"
	"github.com/mavr/anonymous-mail/pkg/config"
	"github.com/mavr/anonymous-mail/pkg/msgrecv/ucmsgrecv"
	"github.com/mavr/anonymous-mail/pkg/msgsnd/ucmsgsnd"
	"github.com/mavr/anonymous-mail/pkg/storage/local"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var (
	version  string
	revision string
)

func main() {
	conf, err := config.New("conf/config.toml")
	if err != nil {
		fmt.Printf("Cannot load config. Error : %s\n", err.Error())
		return
	}

	logConfigure(conf.App.Debug)

	log.Info("Starting backend service for anonymous telegram mailer")
	log.WithField("api_version", version).WithField("revision", revision).Info("Build version")
	if conf.App.Debug {
		log.Info("Running in debug mode")
	}

	ctx, ctxshutdown := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	shutdown := func(err error, force bool) {
		exitCode := 0

		ctxshutdown()

		if !force {
			wg.Wait()
		}

		if err != nil {
			exitCode = 0
			log.WithError(err).Error("Service stopped")
		} else {
			log.Warn("Service stopped")
		}

		os.Exit(exitCode)
	}

	// Initializing local storage
	store, err := local.New()
	if err != nil {
		shutdown(
			errors.Wrap(err, "failed during create local storage"),
			false,
		)
	}

	bot, err := anonbot.New(*conf)
	if err != nil {
		log.WithError(err).Error("Failed bot initializing")

		return
	}
	log.WithField("bot_name", bot.Self().UserName).Info("Bot initialization")

	recv := ucmsgrecv.New(store, bot, ucmsgrecv.Configuration{
		NumberJobs: 2,
		Debug:      conf.App.Debug,
	})

	send := ucmsgsnd.New(store, bot, ucmsgsnd.Configuration{
		Jobs: 1,
	})

	go func() {
		wg.Add(1)
		defer wg.Done()

		if err := recv.Processing(ctx); err != nil {
			log.WithError(err).Error("Receiver failed")
			return
		}

		log.Info("Bot receive process stoping")
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()

		if err := send.Processing(ctx); err != nil {
			log.WithError(err).Error("Sender failed")
			return
		}

		log.Info("Bot sender process stoping")
	}()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	s := <-sigc

	log.WithField("signal", s.String()).Info("Receive os signal")

	shutdown(nil, false)
}

func logConfigure(debug bool) {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006.01.02 15:04:05",
	})

	if debug {
		log.SetLevel(log.DebugLevel)

		return
	}

	log.SetLevel(log.InfoLevel)
}
