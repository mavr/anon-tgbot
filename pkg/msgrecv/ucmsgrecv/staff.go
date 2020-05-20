package ucmsgrecv

import "github.com/mavr/anonymous-mail/pkg/anonbot"

var (
	staffWellcomeMessage = anonbot.StaffMessageCreate(
		anonbot.Translate{Lang: anonbot.LangDefault, Text: `*Wellcome to anonymous bot mailer!*
		
To send anonymous message for user use 
@username text of the message you want to send`},

		anonbot.Translate{Lang: anonbot.LangRus, Text: `*Добро пожаловать в отправляющий анонимные сообщения бот!*
		
Для того что б послать анонимное сообщение человеку используйте следующий формат:
@username текст сообщения, которое вы хотите отправить.
		`},
	)

	staffErrWrongCommand = anonbot.StaffMessageCreate(
		anonbot.Translate{Lang: anonbot.LangDefault, Text: "Command not found."},
		anonbot.Translate{Lang: anonbot.LangRus, Text: "Неверная команда."},
	)

	staffErrNotDeliver = anonbot.StaffMessageCreate(
		anonbot.Translate{Lang: anonbot.LangDefault, Text: "Message could not be delivered."},
		anonbot.Translate{Lang: anonbot.LangRus, Text: "Сообщение не может быть отправленою"},
	)

	staffWrongMessageFormat = anonbot.StaffMessageCreate(
		anonbot.Translate{Lang: anonbot.LangDefault, Text: "Wrong message format."},
		anonbot.Translate{Lang: anonbot.LangRus, Text: "Формат вашего сообщения неверен."},
	)
)
