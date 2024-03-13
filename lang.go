package gnlib

import (
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

var langMap = map[string]string{
	"Afrikaans":  "afr",
	"Arabic":     "ara",
	"Chinese":    "zho",
	"Danish":     "dan",
	"English":    "eng",
	"French":     "fra",
	"German":     "deu",
	"Greek":      "ell",
	"Hausa":      "hau",
	"Hawaiian":   "haw",
	"Indonesian": "ind",
	"Italian":    "ita",
	"Japanese":   "jpn",
	"Korean":     "kor",
	"Malagasy":   "mlg",
	"Portuguese": "por",
	"Romanian":   "ron",
	"Slovenian":  "slv",
	"Spanish":    "spa",
	"Swedish":    "swe",
	"Thai":       "tha",
	"Zulu":       "zul",
}

func LangCode(lang string) string {
	var res string
	tag, err := language.Parse(strings.ToLower(lang))
	if err == nil {
		base, _ := tag.Base()
		res = base.ISO3()
	} else {
		if iso, ok := langMap[lang]; ok {
			res = iso
		}
	}

	return res
}

func LangName(code string) string {
	namer := display.English.Languages()

	tag, err := language.Parse(code)
	if err != nil {
		return ""
	}

	return namer.Name(tag)
}
