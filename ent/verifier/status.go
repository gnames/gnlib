package verifier

type TaxonomicStatus int

const (
	Unknown TaxonomicStatus = iota
	Accepted
	Synonym
)

var txStatusMap = map[string]TaxonomicStatus{
	"not provided": Unknown,
	"accepted":     Accepted,
	"synonym":      Synonym,
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
	return Unknown
}

func (ts TaxonomicStatus) String() string {
	return txStatusStringMap[ts]
}
