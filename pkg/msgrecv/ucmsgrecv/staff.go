package ucmsgrecv

const (
	langDefault = iota
	langRus
)

var ietfToLang = map[string]int{
	"ru": langRus,
}

var (
	staffWellcomeMessage = staffCreate(
		nConstruct{Lang: langDefault, Text: `*Wellcome to anonymouse bot mailer!*
		
		To send anonymouse message for user use 
		@username text of the message you want to send`},

		nConstruct{Lang: langRus, Text: `*Добро пожаловать в отправляющий анонимные сообщения бот!*
		
		Для того что б послать анонимное сообщение человеку используйте следующий формат:
		@username текст сообщения, которое вы хотите отправить.
		`},
	)

	staffErrWrongCommand = staffCreate(
		nConstruct{Lang: langDefault, Text: "Command not found."},
		nConstruct{Lang: langRus, Text: "Неверная команда."},
	)

	staffErrNotDeliver = staffCreate(
		nConstruct{Lang: langDefault, Text: "Message could not be delivered."},
		nConstruct{Lang: langRus, Text: "Сообщение не может быть отправленою"},
	)

	staffWrongMessageFormat = staffCreate(
		nConstruct{Lang: langDefault, Text: "Wrong message format."},
		nConstruct{Lang: langRus, Text: "Формат вашего сообщения неверен."},
	)
)

type nConstruct struct {
	Lang int
	Text string
}

type staff struct {
	m map[int]string
}

func staffCreate(nc ...nConstruct) staff {
	n := staff{
		m: make(map[int]string),
	}

	for _, v := range nc {
		n.m[v.Lang] = v.Text
	}

	return n
}

func (n *staff) Get(languageCode string) string {
	return n.m[ietfToLang[languageCode]]
}
