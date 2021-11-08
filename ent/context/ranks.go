package context

import "strings"

type Rank int

const (
	Empty Rank = iota
	Unknown
	SubSpecies
	Species
	Genus
	Family
	Order
	Class
	Phylum
	Kingdom
)

func (r Rank) String() string {
	return RankStr[r]
}

var RankStr = map[Rank]string{
	Empty:      "empty",
	Unknown:    "unknown",
	SubSpecies: "subspecies",
	Species:    "species",
	Genus:      "genus",
	Family:     "family",
	Order:      "order",
	Class:      "class",
	Phylum:     "phylum",
	Kingdom:    "kingdom",
}

type rankData struct {
	rank  Rank
	total int
	data  map[Clade]int
}

func ranksData() []rankData {
	return []rankData{
		{rank: Kingdom, data: make(map[Clade]int)},
		{rank: Phylum, data: make(map[Clade]int)},
		{rank: Class, data: make(map[Clade]int)},
		{rank: Order, data: make(map[Clade]int)},
		{rank: Family, data: make(map[Clade]int)},
		{rank: Genus, data: make(map[Clade]int)},
		{rank: Species, data: make(map[Clade]int)},
		{rank: SubSpecies, data: make(map[Clade]int)},
		{rank: Unknown, data: make(map[Clade]int)},
		{rank: Empty, data: make(map[Clade]int)},
	}
}

func (r Rank) Index() int {
	i := int(r)
	l := len(RankStr)
	return l - i - 1
}

var StrRank = func() map[string]Rank {
	res := make(map[string]Rank)
	for k, v := range RankStr {
		res[v] = k
	}
	return res
}()

func NewRank(s string) Rank {
	s = strings.ToLower(s)
	if s == "division" {
		s = "phylum"
	}
	if s == "forma" || s == "variety" || s == "ssp" || strings.HasPrefix(s, "subsp") {
		s = "subspecies"
	}
	if rank, ok := StrRank[s]; ok {
		return rank
	}
	return Unknown
}

func AddRank(cs []Clade) {
	for i := range cs {
		if cs[i].Rank == Empty {
			cs[i].Rank = NewRank(cs[i].RankStr)
		}
	}
}
