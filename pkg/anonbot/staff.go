package anonbot

type langCode int

const (
	// LangDefault is english language
	LangDefault langCode = iota

	// LangRus russian
	LangRus
)

var ietfToLang = map[string]langCode{
	"ru": LangRus,
}

// Translate contain translated text and Lang code
type Translate struct {
	Lang langCode
	Text string
}

// Staff struct with translation set
type Staff struct {
	m map[langCode]string
}

// StaffMessageCreate create message and complete this
func StaffMessageCreate(nc ...Translate) Staff {
	n := Staff{
		m: make(map[langCode]string),
	}

	for _, v := range nc {
		n.m[v.Lang] = v.Text
	}

	return n
}

// Get return message with current translate
func (n *Staff) Get(languageCode string) string {
	return n.m[ietfToLang[languageCode]]
}
