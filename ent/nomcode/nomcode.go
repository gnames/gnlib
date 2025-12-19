// Package nomcode provides types and functions for nomenclatural codes.
package nomcode

import "strings"

// Code provides types of nomenclatural Codes.
type Code int

// Constants for different nomenclatural codes.
const (
	Unknown           Code = iota
	Bacterial              // Bacteriological Code
	Botanical              // Botanical Code
	Cultivars              // Cultivated Plant Code
	PhytoSociological      // Phytosociological Code
	Virus                  // Virus Code
	Zoological             // Zoological Code
)

// NewCode converts a string (number or word) to Code.
func New(s string) Code {
	s = strings.ToLower(s)
	switch s {
	case "1", "bact", "bacterial", "icnp":
		return Bacterial
	case "2", "bot", "botanical", "icn", "icnafp", "icbn":
		return Botanical
	case "3", "cult", "cultivar", "cultivars", "icncp":
		return Cultivars
	case "4", "phyto", "phytosociological", "icpn":
		return PhytoSociological
	case "5", "vir", "viral", "virus", "ictv", "icvcn":
		return Virus
	case "6", "zoo", "zoological", "iczn":
		return Zoological
	default:
		return Unknown
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

// String converts code ID to lower case.
func (nc Code) String() string {
	return strings.ToLower(nc.ID())
}

// Abbr returns common abbreviation of the code that is most popular
// in databases and literature.
func (nc Code) Abbr() string {
	switch nc {
	case Bacterial:
		return "ICNP"
	case Botanical:
		return "ICN"
	case Cultivars:
		return "ICNCP"
	case PhytoSociological:
		return "ICPN"
	case Virus:
		return "ICVCN"
	case Zoological:
		return "ICZN"
	default:
		return ""
	}
}
