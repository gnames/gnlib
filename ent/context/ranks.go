package context

import "strings"

type Rank int

const (
	Empty Rank = iota
	Unknown
	SubSpecies
	Species
	SuperSpecies
	SubGenus
	Genus
	SuperGenus
	SubTribe
	Tribe
	InfraFamily
	SubFamily
	Family
	SuperFamily
	InfraOrder
	SubOrder
	Order
	SuperOrder
	ParvClass
	SubTerClass
	InfraClass
	SubClass
	Class
	SuperClass
	SubPhylum
	Phylum
	SuperPhylum
	SubKingdom
	Kingdom
	SuperKingdom
	Empire
)

func (r Rank) String() string {
	return RankStr[r]
}

var RankStr = map[Rank]string{
	Empty:        "empty",
	Unknown:      "unknown",
	SubSpecies:   "subspecies",
	Species:      "species",
	SuperSpecies: "superspecies",
	SubGenus:     "subgenus",
	Genus:        "genus",
	SuperGenus:   "supergenus",
	SubTribe:     "subtribe",
	Tribe:        "tribe",
	InfraFamily:  "infrafamily",
	SubFamily:    "subfamily",
	Family:       "family",
	SuperFamily:  "superfamily",
	InfraOrder:   "infraorder",
	SubOrder:     "suborder",
	Order:        "order",
	SuperOrder:   "superorder",
	ParvClass:    "parvclass",
	SubTerClass:  "subterclass",
	InfraClass:   "infraclass",
	SubClass:     "subclass",
	Class:        "class",
	SuperClass:   "superclass",
	SubPhylum:    "subphylum",
	Phylum:       "phylum",
	SuperPhylum:  "superphylum",
	SubKingdom:   "subkingdom",
	Kingdom:      "kingdom",
	SuperKingdom: "superkingdom",
	Empire:       "empire",
}

type rankData struct {
	rank  Rank
	total int
	data  map[Clade]int
}

func ranksData() []rankData {
	return []rankData{
		{rank: Empire, data: make(map[Clade]int)},
		{rank: SuperKingdom, data: make(map[Clade]int)},
		{rank: Kingdom, data: make(map[Clade]int)},
		{rank: SubKingdom, data: make(map[Clade]int)},
		{rank: SuperPhylum, data: make(map[Clade]int)},
		{rank: Phylum, data: make(map[Clade]int)},
		{rank: SubPhylum, data: make(map[Clade]int)},
		{rank: SuperClass, data: make(map[Clade]int)},
		{rank: Class, data: make(map[Clade]int)},
		{rank: SubClass, data: make(map[Clade]int)},
		{rank: InfraClass, data: make(map[Clade]int)},
		{rank: SubTerClass, data: make(map[Clade]int)},
		{rank: ParvClass, data: make(map[Clade]int)},
		{rank: SuperOrder, data: make(map[Clade]int)},
		{rank: Order, data: make(map[Clade]int)},
		{rank: SubOrder, data: make(map[Clade]int)},
		{rank: InfraOrder, data: make(map[Clade]int)},
		{rank: SuperFamily, data: make(map[Clade]int)},
		{rank: Family, data: make(map[Clade]int)},
		{rank: SubFamily, data: make(map[Clade]int)},
		{rank: InfraFamily, data: make(map[Clade]int)},
		{rank: Tribe, data: make(map[Clade]int)},
		{rank: SubTribe, data: make(map[Clade]int)},
		{rank: SuperGenus, data: make(map[Clade]int)},
		{rank: Genus, data: make(map[Clade]int)},
		{rank: SubGenus, data: make(map[Clade]int)},
		{rank: SuperSpecies, data: make(map[Clade]int)},
		{rank: Species, data: make(map[Clade]int)},
		{rank: SubSpecies, data: make(map[Clade]int)},
		{rank: Unknown, data: make(map[Clade]int)},
		{rank: Empty, data: make(map[Clade]int)},
	}
}

// Index returns index of the rank position in the ranksData.
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
