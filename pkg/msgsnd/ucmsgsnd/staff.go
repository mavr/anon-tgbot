package ucmsgsnd

import "github.com/mavr/anonymous-mail/pkg/anonbot"

var (
	staffNewMessageTitle = anonbot.StaffMessageCreate(
		anonbot.Translate{Lang: anonbot.LangDefault, Text: "Someone send message for you."},
		anonbot.Translate{Lang: anonbot.LangRus, Text: "Вам адресовано новое сообщение."},
	)
)
