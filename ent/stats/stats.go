// package stats calculates some statistics for a group of scientific
// names.
//
// It uses data received from verification of the names against the
// Catalogue of Life, finds distribution of names across Kingdoms and
// finds a taxon that contains a given percentage (always a majority)
// of scientific names of genera and lower.
package stats

// Taxon struct represents a particular taxon according to the Catalogue of
// Life (CoL). It includes an ID from CoL, name of the taxon, and numerical and
// string representation of the taxon's rank.
type Taxon struct {
	// ID is the Catalogue of Life ID for the taxon.
	ID string

	// Name is the name of the taxon.
	Name string

	// RankStr is a string representation of the taxon's rank.
	RankStr string

	// Rank represents taxon's rank via Rank type. Rank type is derived from
	// int type.
	Rank
}

// Stats struct provides statistical data about a group of verified by the
// Catalogue of Life scientific names. It contains data about names number
// used for the stats calculation, the distribution of these names across
// Kingdoms registered in CoL, as well as the lowest taxon that contains
// at least a majority of all these names. A user submits the desired
// threshold for the calculation of such taxon.
type Stats struct {
	// NamesNum is the number of names that are used for stats calculation.
	// These names include names of a rank `genus` and lower,
	// verified to the Catalogue of Life
	NamesNum int

	//Kingdom is the most prevalent kingdom in the group of names.
	Kingdom Taxon

	// MainTaxon is the taxon that contains at least the percentage of names
	// according to the MainTaxonThreshold
	MainTaxon Taxon

	// KingdomPercentage is a value between 0 and 1 representing the percentage
	// of names located in a particular kingdom.
	KingdomPercentage float32

	// MainTaxonPercentage is a value between 0 and 1 representing the
	// percentage of names located in the MainTaxon.
	MainTaxonPercentage float32

	// Kingdoms is the distribution of names across detected kingdoms.
	Kingdoms []TaxonDist
}

// TaxonDist provides information how a group of names is distributed
// across taxons of the same rank.
type TaxonDist struct {
	// NamesNum is the number of names found for this particular rank.
	NamesNum int

	// Name is the scientific name of the taxon.
	Name string

	// Percentage is the percentage of names belonging to this taxon.
	Percentage float32
}

// New takes several hierarhies, a MainTaxon threshold value, and returns back
// the kingdom where most of items belong to (if rank 'kingdom' is provided),
// percentage of how many items belong to that kingdom, and the highest ranking
// taxon that includes at least the given percentage of species. The percentage
// is provided via threshold parameter.
//
// The algorithm assumes that all items belong to the same classification tree.
func New(
	h []Hierarchy,
	threshold float32,
) Stats {
	if threshold < 0.5 {
		threshold = 0.5
	}

	// collect names that are genus or lower, no taxons are removed from
	// the hierarchy.
	taxons := extractTaxons(h)
	if len(taxons) == 1 {
		return Stats{}
	}
	namesNum := len(taxons)

	// get empty structure for ranks stats
	ranks := ranksData()
	// populate ranks
	for _, cs := range taxons {
		for i := range cs {
			rankIdx := cs[i].Index()
			ranks[rankIdx].data[cs[i]]++
			ranks[rankIdx].total++
		}
	}

	ranks = removeEmptyRanks(ranks)
	res := calcStats(namesNum, ranks, threshold)
	return res
}

func calcStats(
	namesNum int,
	ranks []rankData,
	threshold float32,
) Stats {
	var ks []TaxonDist
	var kingdom, mainTaxon Taxon
	var kPCent, txnPCent float32
	var foundMainTaxon bool
	l := len(ranks)

	for i := range ranks {
		revI := l - 1 - i
		if ranks[revI].rank <= Unknown {
			continue
		}
		txn, pcent := maxTaxon(namesNum, ranks[revI])
		if ranks[revI].rank == Kingdom {
			ks = getKingdoms(ranks[revI])
			if isMaxKingdom(ks, pcent) {
				kingdom, kPCent = txn, pcent
			}
		}
		if pcent > threshold && !foundMainTaxon {
			mainTaxon, txnPCent = txn, pcent
			foundMainTaxon = true
		}
	}
	return Stats{
		NamesNum:            namesNum,
		Kingdom:             kingdom,
		MainTaxon:           mainTaxon,
		KingdomPercentage:   kPCent,
		MainTaxonPercentage: txnPCent,
		Kingdoms:            ks,
	}
}

func isMaxKingdom(cd []TaxonDist, percentage float32) bool {
	var count int
	for i := range cd {
		if cd[i].Percentage == percentage {
			count++
		}
	}
	return count == 1
}

func getKingdoms(kingdom rankData) []TaxonDist {
	res := make([]TaxonDist, len(kingdom.data))
	var i int
	for k, v := range kingdom.data {
		cd := TaxonDist{
			NamesNum:   v,
			Name:       k.Name,
			Percentage: float32(v) / float32(kingdom.total),
		}
		res[i] = cd
		i++
	}
	return res
}

func maxTaxon(namesNum int, rd rankData) (Taxon, float32) {
	var max int
	var res, cld Taxon
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

// extractTaxons collects taxons for each name. It only collects names that
// are genus or less. It does not make sense to take in account higher
// classification ranks because their meaning can be different than in
// the Catalogue of Life.
func extractTaxons(h []Hierarchy) [][]Taxon {
	var taxons []Taxon
	res := make([][]Taxon, 0, len(h))
	for i := range h {
		var genusOrLess bool
		taxons = h[i].Taxons()
		for ii := range taxons {
			if taxons[ii].Rank == Empty {
				taxons[ii].Rank = NewRank(taxons[ii].RankStr)
			}
			if !genusOrLess &&
				taxons[ii].Rank != Unknown &&
				taxons[ii].Rank <= Genus {
				genusOrLess = true
			}
		}
		if genusOrLess {
			res = append(res, taxons)
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
