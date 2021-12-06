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

	clades := extractClades(h)
	if len(clades) == 1 {
		return Context{}
	}

	ranks := ranksData()
	for _, cs := range clades {
		for i := range cs {
			rankIdx := cs[i].Index()
			ranks[rankIdx].data[cs[i]]++
			ranks[rankIdx].total++
		}
	}
	ranks = cleanupRanks(ranks)
	return calcContext(ranks, threshold)
}

func calcContext(ranks []rankData,
	threshold float32,
) Context {
	var ks []CladesDist
	var kingdom, context Clade
	var kPC, cPC float32

	for i := range ranks {
		c, pc := maxClade(ranks[i])
		if ranks[i].rank == Kingdom {
			ks = getKingdoms(ranks[i])
			if isMaxKingdom(ks, pc) {
				kingdom, kPC = c, pc
			}
		}
		if pc < threshold {
			if ranks[i].rank < Kingdom {
				break
			} else {
				continue
			}
		}
		context, cPC = c, pc
	}
	return Context{
		Kingdom:           kingdom,
		Context:           context,
		KingdomPercentage: kPC,
		ContextPercentage: cPC,
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

func maxClade(rd rankData) (Clade, float32) {
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
	return res, float32(max) / float32(rd.total)
}

// extractClades removes clades that do not contain all records.
func extractClades(h []Hierarch) [][]Clade {
	res := make([][]Clade, len(h))
	for i := range h {
		res[i] = h[i].Clades()
		for ii := range res[i] {
			if res[i][ii].Rank == Empty {
				res[i][ii].Rank = NewRank(res[i][ii].RankStr)
			}
		}
	}
	return res
}

// cleanupRanks leaves only ranks that are known to context
func cleanupRanks(ranks []rankData) []rankData {
	var res []rankData
	var total int
	for i := range ranks {
		if ranks[i].total == 0 {
			continue
		}
		if total == 0 {
			total = ranks[i].total
		}
		if ranks[i].total == total {
			res = append(res, ranks[i])
		}
	}
	return res
}
