package gnlib_test

import (
	"testing"

	"github.com/gnames/gnlib"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	msg, in, out, lang string
}{
	{"cal", "cal", "cal", ""},
	{"SMO", "SMO", "smo", "Samoan"},
	{"nob", "nob", "nob", "Norwegian Bokmål"},
	{"@zh", "@zh", "", ""},
	{"CZE", "CZE", "ces", "Czech"},
	{"new", "new", "new", "Newari"},
	{"nzi", "nzi", "nzi", "Nzima"},
	{"fin", "fin", "fin", "Finnish"},
	{"min", "min", "min", "Minangkabau"},
	{"gcf", "gcf", "gcf", ""},
	{"aa", "aa", "aar", "Afar"},
	{"skg", "skg", "skg", ""},
	{"tr", "tr", "tur", "Turkish"},
	{"CHA", "CHA", "cha", "Chamorro"},
	{"agn", "agn", "agn", ""},
	{"ga", "ga", "gle", "Irish"},
	{"kha", "kha", "kha", "Khasi"},
	{"hts", "hts", "hts", ""},
	{"mug", "mug", "mug", ""},
	{"inn", "inn", "inn", ""},
	{"ee", "ee", "ewe", "Ewe"},
	{"BER", "BER", "ber", ""},
	{"pms", "pms", "pms", "Piedmontese"},
	{"yap", "yap", "yap", "Yapese"},
	{"nez", "nez", "nez", ""},
	{"fub", "fub", "fub", ""},
	{"pon", "pon", "pon", "Pohnpeian"},
	{"usa", "usa", "usa", ""},
	{"vep", "vep", "vep", "Veps"},
	{"paa", "paa", "paa", ""},
	{"fy", "fy", "fry", "Western Frisian"},
	{"guj", "guj", "guj", "Gujarati"},
	{"ppl", "ppl", "ppl", ""},
	{"dan", "dan", "dan", "Danish"},
	{"TUR", "TUR", "tur", "Turkish"},
	{"qug", "qug", "qug", "Chimborazo Highland Quichua"},
	{"nan", "nan", "nan", "Min Nan Chinese"},
	{"oci", "oci", "oci", "Occitan"},
	{"lo", "lo", "lao", "Lao"},
	{"veo", "veo", "veo", ""},
	{"fij", "fij", "fij", "Fijian"},
	{"cr", "cr", "cre", "Cree"},
	{"nl", "nl", "nld", "Dutch"},
	{"cym", "cym", "cym", "Welsh"},
	{"gu", "gu", "guj", "Gujarati"},
	{"PAP", "PAP", "pap", "Papiamento"},
	{"sun", "sun", "sun", "Sundanese"},
	{"Australia Aboriginal", "Australia Aboriginal", "", ""},
	{"RUS", "RUS", "rus", "Russian"},
	{"SLK", "SLK", "slk", "Slovak"},
	{"sin", "sin", "sin", "Sinhala"},
	{"xh", "xh", "xho", "Xhosa"},
	{"sah", "sah", "sah", "Sakha"},
	{"loo", "loo", "loo", ""},
	{"spa", "spa", "spa", "Spanish"},
	{"cjk", "cjk", "cjk", ""},
	{"KOR", "KOR", "kor", "Korean"},
	{"TGL", "TGL", "fil", "Filipino"},
	{"hu", "hu", "hun", "Hungarian"},
	{"st", "st", "sot", "Southern Sotho"},
	{"frr", "frr", "frr", "Northern Frisian"},
	{"WOL", "WOL", "wol", "Wolof"},
	{"tui", "tui", "tui", ""},
	{"FChinaEng", "FChinaEng", "", ""},
	{"vls", "vls", "vls", "West Flemish"},
	{"tem", "tem", "tem", "Timne"},
	{"azb", "azb", "azb", ""},
	{"oj", "oj", "oji", "Ojibwa"},
	{"eet", "eet", "", ""},
	{"loz", "loz", "loz", "Lozi"},
	{"ell", "ell", "ell", "Greek"},
	{"nyf", "nyf", "nyf", ""},
	{"GLA", "GLA", "gla", "Scottish Gaelic"},
	{"bjs", "bjs", "bjs", ""},
	{"ARM", "ARM", "hye", "Armenian"},
	{"bsy", "bsy", "bsy", ""},
	{"kal", "kal", "kal", "Kalaallisut"},
}

func TestLangCode(t *testing.T) {
	assert := assert.New(t)
	for _, v := range tests {
		code := gnlib.LangCode(v.in)
		lang := gnlib.LangName(code)
		assert.Equal(v.out, code, v.msg)
		assert.Equal(v.lang, lang, v.msg)
	}
}
