package context

type Clade struct {
	ID, Name, RankStr string
	Rank
}

type Context struct {
	NamesNum                             int
	Kingdom, Context                     Clade
	KingdomPercentage, ContextPercentage float32
	Kingdoms                             []CladesDist
}

type CladesDist struct {
	NamesNum   int
	Name       string
	Percentage float32
}

// New takes several items that include bio-clasification and returns back
// the kingdom where most of items belong to (if rank 'kingdom' is provided),
// percentage of how many items belong to that kingdom, and the highest ranking
// clade that includes a certain percentage of species. The percentage is
// provided via threshold parameter.
//
// The algorithm assumes that all items belong to the same classification tree
// and that classification does not skip clades from item to item.
func New(
	h []Hierarch,
	threshold float32,
) Context {
	if threshold <= 0.5 {
		threshold = 0.5001
	}

	// collect names that are genus or less, no clades are removed from
	// the hierarchy.
	clades := extractClades(h)
	if len(clades) == 1 {
		return Context{}
	}
	namesNum := len(clades)

	// get empty structure for ranks stats
	ranks := ranksData()
	// populate ranks
	for _, cs := range clades {
		for i := range cs {
			rankIdx := cs[i].Index()
			ranks[rankIdx].data[cs[i]]++
			ranks[rankIdx].total++
		}
	}

	ranks = removeEmptyRanks(ranks)
	res := calcContext(namesNum, ranks, threshold)
	return res
}

func calcContext(
	namesNum int,
	ranks []rankData,
	threshold float32,
) Context {
	var ks []CladesDist
	var kingdom, context Clade
	var kPCent, cPCent float32

	for i := range ranks {
		if ranks[i].rank <= Unknown {
			break
		}
		c, pcent := maxClade(namesNum, ranks[i])
		if ranks[i].rank == Kingdom {
			ks = getKingdoms(ranks[i])
			if isMaxKingdom(ks, pcent) {
				kingdom, kPCent = c, pcent
			}
		}
		if pcent < threshold {
			if ranks[i].rank < Kingdom {
				break
			} else {
				continue
			}
		}
		context, cPCent = c, pcent
	}
	return Context{
		NamesNum:          namesNum,
		Kingdom:           kingdom,
		Context:           context,
		KingdomPercentage: kPCent,
		ContextPercentage: cPCent,
		Kingdoms:          ks,
	}
}

func isMaxKingdom(cd []CladesDist, percentage float32) bool {
	var count int
	for i := range cd {
		if cd[i].Percentage == percentage {
			count++
		}
	}
	return count == 1
}

func getKingdoms(kingdom rankData) []CladesDist {
	res := make([]CladesDist, len(kingdom.data))
	var i int
	for k, v := range kingdom.data {
		cd := CladesDist{
			NamesNum:   v,
			Name:       k.Name,
			Percentage: float32(v) / float32(kingdom.total),
		}
		res[i] = cd
		i++
	}
	return res
}

func maxClade(namesNum int, rd rankData) (Clade, float32) {
	var max int
	var res, cld Clade
	for k, v := range rd.data {
		if v > max {
			max = v
			cld = k
		}
	}
	if cld.Name != "" {
		res = cld
	}
	return res, float32(max) / float32(namesNum)
}

// extractClades collects clades for each name. It only collects names that
// are genus or less. It does not make sense to take in account higher
// classification ranks because their meaning can be different than in
// the Catalogue of Life.
func extractClades(h []Hierarch) [][]Clade {
	var clades []Clade
	res := make([][]Clade, 0, len(h))
	for i := range h {
		var genusOrLess bool
		clades = h[i].Clades()
		for ii := range clades {
			if clades[ii].Rank == Empty {
				clades[ii].Rank = NewRank(clades[ii].RankStr)
			}
			if !genusOrLess &&
				clades[ii].Rank != Unknown &&
				clades[ii].Rank <= Genus {
				genusOrLess = true
			}
		}
		if genusOrLess {
			res = append(res, clades)
		}
	}
	return res
}

// removeEmptyRanks removes empty ranks
func removeEmptyRanks(ranks []rankData) []rankData {
	var res []rankData
	for i := range ranks {
		if ranks[i].total == 0 {
			continue
		}
		res = append(res, ranks[i])
	}
	return res
}
