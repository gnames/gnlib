package verifier

type TaxonomicStatus int

const (
	UnknownTaxStatus TaxonomicStatus = iota
	AcceptedTaxStatus
	SynonymTaxStatus
)

var txStatusMap = map[string]TaxonomicStatus{
	"not provided": UnknownTaxStatus,
	"accepted":     AcceptedTaxStatus,
	"synonym":      SynonymTaxStatus,
}

var txStatusStringMap = func() map[TaxonomicStatus]string {
	res := make(map[TaxonomicStatus]string)
	for k, v := range txStatusMap {
		res[v] = k
	}
	return res
}()

func New(txStatus string) TaxonomicStatus {
	if res, ok := txStatusMap[txStatus]; ok {
		return res
	}
	return UnknownTaxStatus
}

func (ts TaxonomicStatus) String() string {
	return txStatusStringMap[ts]
}
