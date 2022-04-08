package stats

import "strings"

// Rank represents a rank of a taxon.
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

// String returns the string representation of a Rank.
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
	data  map[Taxon]int
}

func ranksData() []rankData {
	return []rankData{
		{rank: Empire, data: make(map[Taxon]int)},
		{rank: SuperKingdom, data: make(map[Taxon]int)},
		{rank: Kingdom, data: make(map[Taxon]int)},
		{rank: SubKingdom, data: make(map[Taxon]int)},
		{rank: SuperPhylum, data: make(map[Taxon]int)},
		{rank: Phylum, data: make(map[Taxon]int)},
		{rank: SubPhylum, data: make(map[Taxon]int)},
		{rank: SuperClass, data: make(map[Taxon]int)},
		{rank: Class, data: make(map[Taxon]int)},
		{rank: SubClass, data: make(map[Taxon]int)},
		{rank: InfraClass, data: make(map[Taxon]int)},
		{rank: SubTerClass, data: make(map[Taxon]int)},
		{rank: ParvClass, data: make(map[Taxon]int)},
		{rank: SuperOrder, data: make(map[Taxon]int)},
		{rank: Order, data: make(map[Taxon]int)},
		{rank: SubOrder, data: make(map[Taxon]int)},
		{rank: InfraOrder, data: make(map[Taxon]int)},
		{rank: SuperFamily, data: make(map[Taxon]int)},
		{rank: Family, data: make(map[Taxon]int)},
		{rank: SubFamily, data: make(map[Taxon]int)},
		{rank: InfraFamily, data: make(map[Taxon]int)},
		{rank: Tribe, data: make(map[Taxon]int)},
		{rank: SubTribe, data: make(map[Taxon]int)},
		{rank: SuperGenus, data: make(map[Taxon]int)},
		{rank: Genus, data: make(map[Taxon]int)},
		{rank: SubGenus, data: make(map[Taxon]int)},
		{rank: SuperSpecies, data: make(map[Taxon]int)},
		{rank: Species, data: make(map[Taxon]int)},
		{rank: SubSpecies, data: make(map[Taxon]int)},
		{rank: Unknown, data: make(map[Taxon]int)},
		{rank: Empty, data: make(map[Taxon]int)},
	}
}

// Index returns the index of a rank position in the ranksData.
func (r Rank) Index() int {
	i := int(r)
	l := len(RankStr)
	return l - i - 1
}

// StrRank conversts a rank string to Rank type.
var StrRank = func() map[string]Rank {
	res := make(map[string]Rank)
	for k, v := range RankStr {
		res[v] = k
	}
	return res
}()

// NewRank creates Rank from a string.
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

// AddRank converts a RankStr to its Rank value and saves it in taxons.
func AddRank(cs []Taxon) {
	for i := range cs {
		if cs[i].Rank == Empty {
			cs[i].Rank = NewRank(cs[i].RankStr)
		}
	}
}
