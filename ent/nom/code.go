package nom

import "strings"

// Code provides types of nomenclatural Codes.
type Code int

// Constants for different nomenclatural codes.
const (
	UnknownCode       Code = iota
	Bacterial              // Bacteriological Code
	Botanical              // Botanical Code
	Cultivars              // Cultivated Plant Code
	PhytoSociological      // Phytosociological Code
	Virus                  // Virus Code
	Zoological             // Zoological Code
)

// NewCode converts a string (number or word) to Code.
func NewCode(s string) Code {
	s = strings.ToLower(s)
	switch s {
	case "1", "bacterial", "icnp":
		return Bacterial
	case "2", "botanical", "icn", "icnafp", "icbn":
		return Botanical
	case "3", "cultivars", "icncp":
		return Cultivars
	case "4", "phytosociological", "icpn":
		return PhytoSociological
	case "5", "virus", "icvcn":
		return Virus
	case "6", "zoological", "iczn":
		return Zoological
	default:
		return UnknownCode
	}
}

var CodeToString = map[Code]string{
	Bacterial:         "BACTERIAL",
	Botanical:         "BOTANICAL",
	Cultivars:         "CULTIVARS",
	PhytoSociological: "PHYTOSOCIOLOGICAL",
	Virus:             "VIRUS",
	Zoological:        "ZOOLOGICAL",
}

// ID returns capitalized name of the code. It corresponds to ID in
// SFGA schema.
func (nc Code) ID() string {
	if res, ok := CodeToString[nc]; ok {
		return res
	}
	return ""
}

func (nc Code) String() string {
	return ToStr(nc.ID())
}

// ToStr normalizes enumerated string IDs to 'normal' strings.
// For example 'PROVISIONALLY_ACCEPTED' becomes
// 'provisionally accepted'.
func ToStr(s string) string {
	s = strings.ToLower(s)
	return strings.ReplaceAll(s, "_", " ")
}
