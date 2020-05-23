package apiserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var startAt time.Time

const versionAPI = "0.1"

func init() {
	startAt = time.Now()
}

// APIServiceConfig configuration struct for api service
type APIServiceConfig struct {
	AppRevision  string
	AppVersion   string
	AppDebugMode bool

	ServPort int
}
type apiService struct {
	s    *http.Server
	port int
}

// NewAPIService create new api service
func NewAPIService(c APIServiceConfig) *apiService {
	r := gin.Default()

	if !c.AppDebugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	apiGR := r.Group("/api")
	apiGR.GET("/status", statusHandle(c.AppVersion, c.AppRevision))

	return &apiService{
		s: &http.Server{
			Handler: r,
			Addr:    fmt.Sprintf(":%d", c.ServPort),
		},
	}
}

func (s *apiService) Run(ctx context.Context) error {
	go func(ctx context.Context) {
		<-ctx.Done()
		ctxShutdown, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		if err := s.s.Shutdown(ctxShutdown); err != nil {
			log.WithError(err).Error("api server shutdown failed")
		}
	}(ctx)

	if err := s.s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func statusHandle(version, revision string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, StatusResponse{
			Build:      revision,
			Version:    version,
			VersionAPI: versionAPI,
			Uptime:     time.Now().Sub(startAt).Round(time.Second).String(),
		})
	}
}
