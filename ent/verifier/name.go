package verifier

import (
	"strings"

	"github.com/gnames/gnstats/ent/stats"
)

// Name is a result of verification of one name-string from the input.
type Name struct {
	// ID is a UUIDv5 generated out of the Input string.
	ID string `json:"id"`

	// Name is a verified name-string
	Name string `json:"name"`

	// Cardinality is the cardinality of input name:
	// 0 - No match, virus or hybrid formula,
	// 1 - Uninomial, 2 - Binomial, 3 - Trinomial etc.
	Cardinality int `json:"cardinality"`

	// MatchType is best available match.
	MatchType MatchTypeValue `json:"matchType"`

	// BestResult is the best result according to GNames scoring.
	BestResult *ResultData `json:"bestResult,omitempty"`

	// Results contain all detected matches from preverred data sources
	// provided by user.
	Results []*ResultData `json:"results,omitempty"`

	// DataSourcesNum is a number of data sources that matched an
	// input name-string.
	DataSourcesNum int `json:"dataSourcesNum,omitempty"`

	DataSourcesIDs []int `json:"dataSourcesIds,omitempty"`

	// DataSourcesDetails contains information about matched data-sources
	// and the IDs of records that did match.
	DataSourcesDetails []DataSourceDetails `json:"dataSourcesDetails,omitempty"`

	// Curation estimates reliability of matched data sources. If
	// matches are returned by at least one manually curated data source, or by
	// automatically curated data source, or only by sources that are not
	// significantly manually curated.
	Curation CurationLevel `json:"curation"`

	// OverloadDetected might be triggered if a virus name or a canonical name
	// contain many variations and/or strains. In this case not all data are
	// queried.
	OverloadDetected string `json:"overloadDetected,omitempty"`

	// Error provides an error message, if any. If error is not empty, the match
	// failed because of a bug in the service.
	Error string `json:"error,omitempty"`
}

// DataSourceDetails describe data-source and found best match.
type DataSourceDetails struct {
	// DataSourceID is the ID of the DataSource in GNverifier.
	DataSourceID int `json:"dataSourceId"`

	// TitleShort is the short name of the resource.
	TitleShort string `json:"title"`

	// Match is the collection of found records for the data-source.
	Match MatchShort `json:"match"`
}

// MatchShort contains data of a matched record for a resource.
type MatchShort struct {
	// RecordID is the RecordID of a name-string.
	RecordID string `json:"recordId,omitempty"`

	// Outlink is the link to a match in the original data-source.
	Outlink string `json:"outlink,omitempty"`

	// NameString is verbatim representation of a name.
	NameString string `json:"nameString"`

	// AuthScore provides information about AuthorScore from sorting algorithm.
	AuthScore bool `json:"-"`
}

// ResultData are returned data of the `BestResult` or `Results` of
// name verification.
type ResultData struct {
	// DataSourceID is the ID of a matched DataSource.
	DataSourceID int `json:"dataSourceId"`

	// Shortened/abbreviated title of the data source.
	DataSourceTitleShort string `json:"dataSourceTitleShort"`

	// Curation of the data source.
	Curation CurationLevel `json:"curation"`

	// RecordID from a data source. We try our best to return ID that
	// corresponds to dwc:taxonID of a DataSource. If such ID is not provided,
	// this ID will be auto-generated.  Auto-generated IDs will
	// have 'gn_' prefix.
	RecordID string `json:"recordId"`

	// GlobalID that is exposed globally by a DataSource. Such IDs are usually
	// self-resolved, like for example LSID, pURL, DOI etc.
	GlobalID string `json:"globalId,omitempty"`

	// LocalID used by a DataSource internally. If an OutLink field is provided,
	// LocalID serves as a 'dynamic' component of the URL.
	LocalID string `json:"localId,omitempty"`

	// Outlink to the record in the DataSource. It consists of a 'stable'
	// URL and an appended 'dynamic' LocalID
	Outlink string `json:"outlink,omitempty"`

	// EntryDate is a timestamp created on entry of the data.
	EntryDate string `json:"entryDate"`

	// SortScore is a numeric representation of the whole score.
	// It can be used to find the BestMatch overall, as well as the
	// best match for every data-source.
	//
	// SortScore takes data from all other scores, using the priority
	// sequence from highest to lowest: InfraSpecificRankScore, FuzzyLessScore,
	// CuratedDataScore, AuthorMatchScore, AcceptedNameScore,
	// ParsingQualityScore. Every highest priority trumps everything below.
	// When the final score value is calculated, it is used to
	// sort verification or search results.
	//
	// Comparing this score between results of different verifications will
	// not necessary be accurate. The score is used for comparison of names
	// from the same result.
	SortScore float64 `json:"sortScore"`

	// ParsingQuality determines how well gnparser was able to break the
	// name-string to its components. 0 - no parse, 1 - clean parse,
	// 2 - some problems, 3 - significant problems.
	ParsingQuality int `json:"-"`

	// MatchedID is the UUID v5 derived from the MatchedName.
	MatchedNameID string `json:"matchedNameID"`

	// MatchedName is a name-string from the DataSource that was matched
	// by GNames algorithm.
	MatchedName string `json:"matchedName"`

	// MatchCardinality is the cardinality of returned name:
	// 0 - No match, virus or hybrid formula,
	// 1 - Uninomial, 2 - Binomial, 3 - trinomial etc.
	MatchedCardinality int `json:"matchedCardinality"`

	// MatchedCanonicalSimple is a simplified canonical form without ranks for
	// names lower than species, and with omitted hybrid signs for named
	// hybrids. Quite often simple canonical is the same as full canonical.
	// Hybrid signs are preserved for hybrid formulas.
	MatchedCanonicalSimple string `json:"matchedCanonicalSimple,omitempty"`

	// MatchedCanonicalFull is a canonical form that preserves hybrid signs
	// and infraspecific ranks.
	MatchedCanonicalFull string `json:"matchedCanonicalFull,omitempty"`

	// MatchedAuthors is a list of authors mentioned in the name.
	MatchedAuthors []string `json:"-"`

	// MatchedYear is a year mentioned in the name. Multiple years or
	// approximate years are ignored.
	MatchedYear int `json:"-"`

	// CurrentRecordID is the id of currently accepted name given by
	// the data-source.
	CurrentRecordID string `json:"currentRecordId"`

	// CurrentID is the UUID v5 derived from the CurrentName.
	CurrentNameID string `json:"currentNameId"`

	// CurrentName is a currently accepted name (it is only provided by
	// DataSources with taxonomic data).
	CurrentName string `json:"currentName"`

	// CurrentCardinality is a cardinality of the accepted name.
	// It might differ from the matched name cardinality.
	CurrentCardinality int `json:"currentCardinality"`

	// CurrentCanonicalSimple is a canonical form for the currently accepted name.
	CurrentCanonicalSimple string `json:"currentCanonicalSimple"`

	// CurrentCanonicalFull is a full version of canonicall form for the
	// currently accepted name.
	CurrentCanonicalFull string `json:"currentCanonicalFull"`

	// TaxonomicStatus provides taxonomic status of a name.
	// Can be "accepted", "synonym", "not provided".
	TaxonomicStatus string `json:"taxonomicStatus"`

	// IsSynonym is a boolean value that is true if the matched name is a
	// synonym. DEPRECATED: use TaxonomicStatus instead.
	IsSynonym bool `json:"isSynonym,omitempty"`

	// ClassificationPath to the name (if provided by the DataSource).
	// Classification path consists of a hierarchy of name-strings.
	ClassificationPath string `json:"classificationPath,omitempty"`

	// ClassificationRanks of the classification path. They follow the
	// same order as the classification path.
	ClassificationRanks string `json:"classificationRanks,omitempty"`

	// ClassificationIDs of the names-strings. They always correspond to
	// the "id" field.
	ClassificationIDs string `json:"classificationIds,omitempty"`

	// EditDistance is a Levenshtein edit distance between canonical form of the
	// input name-string and the matched canonical form. If match type is
	// "EXACT", edit-distance will be 0.
	EditDistance int `json:"editDistance"`

	// StemEditDistance is a Levenshtein edit distance after removing suffixes
	// from specific epithets from canonical forms.
	StemEditDistance int `json:"stemEditDistance"`

	//MatchType describes what kind of a match happened to a name-string.
	MatchType MatchTypeValue `json:"matchType"`

	// ScoreDetails provides data about matching of authors, year, rank,
	// parsingQuality...
	ScoreDetails `json:"scoreDetails"`

	// Vernacular names that correspond to the matched name. (Will be implemented
	// later)
	Vernaculars []Vernacular `json:"vernaculars,omitempty"`
}

// Vernacular name
type Vernacular struct {
	Name string `json:"name"`

	// Language of the name, hopefully in ISO form.
	Language string `json:"language,omitempty"`

	// Locality is geographic places where the name is used.
	Locality string `json:"locality,omitempty"`
}

// ScoreDetails provides explanations how sorting of result occures and
// why something became selected as the `BestResult`. Score data for every
// item is normalized to a range from 0 to 1 where 0 means there were no
// match by the factor, and 1 means a "perfect" match by the item.
// Fields located higher on the list have more weight than lower fields.
// It means that lower fields are getting into account only if higher fields
// provide equal values.
// For all scores 1 is the best, 0 is the worst.
type ScoreDetails struct {
	// CardinalityScore is 1 when cardinality of input name and match name
	// match and neither cardinality is 0. In all other cases it this score
	// equal 0.
	CardinalityScore float32 `json:"cardinalityScore"`

	// InfraSpecificRankScore matches infraspecific rank. For example if a
	// query name is `Aus bus var. cus`, and the match has the same rank,
	// this field is 1.
	InfraSpecificRankScore float32 `json:"infraSpecificRankScore"`

	// FuzzyLessScore scores edit distance for fuzzy matching. If edit distance
	// is 0 the score is maxed to 1.
	FuzzyLessScore float32 `json:"fuzzyLessScore"`

	// CuratedDataScore scores highest if the matched data-source is known for
	// having a significant manual curation effort of the data.
	CuratedDataScore float32 `json:"curatedDataScore"`

	// AuthorMatchScore tries to match authors and years in the name. If
	// a year and all authors match, the score is 1.
	AuthorMatchScore float32 `json:"authorMatchScore"`

	// AcceptedNameScore is a binary field, if matched name is also currently
	// accepted name according to the data-source, the value is 1.
	AcceptedNameScore float32 `json:"acceptedNameScore"`

	// ParsingQualityScore is the highest for matched names that were parsed
	// without any problems.
	ParsingQualityScore float32 `json:"parsingQualityScore"`
}

func (n Name) Taxons() []stats.Taxon {
	var res []stats.Taxon
	if n.BestResult == nil || n.BestResult.DataSourceID != 1 {
		return res
	}

	path := strings.Split(n.BestResult.ClassificationPath, "|")
	ids := strings.Split(n.BestResult.ClassificationIDs, "|")
	ranks := strings.Split(n.BestResult.ClassificationRanks, "|")
	if len(path) < 2 {
		return res
	}

	res = make([]stats.Taxon, len(path))

	for i := range path {
		res[i] = stats.Taxon{Name: path[i]}
		if len(ids) == len(path) {
			res[i].ID = ids[i]
			res[i].ID = ids[i]
		}
		if len(ranks) == len(path) {
			res[i].RankStr = ranks[i]
		}
	}
	return res
}
