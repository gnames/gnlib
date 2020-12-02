package verifier_test

import (
	"testing"
	"time"

	ver "github.com/gnames/gnlib/domain/entity/verifier"
	"github.com/gnames/gnlib/encode"
	"github.com/stretchr/testify/assert"
)

func TestDataSourceJSON(t *testing.T) {
	enc := encode.GNjson{Pretty: true}
	updated, _ := time.Parse("2006-01-02", "2020-06-30")
	testData := []struct {
		ds   ver.DataSource
		json string
	}{
		{
			ds: ver.DataSource{
				ID:          1,
				Title:       "Catalogue of Life",
				TitleShort:  "Catalogue of Life",
				Curation:    ver.Curated,
				RecordCount: 4_000_000,
				UpdatedAt:   updated,
			},
			json: `{
  "id": 1,
  "title": "Catalogue of Life",
  "titleShort": "Catalogue of Life",
  "curation": "Curated",
  "recordCount": 4000000,
  "updatedAt": "2020-06-30T00:00:00Z"
}`,
		},
	}
	for _, v := range testData {
		ser, err := enc.Encode(v.ds)
		assert.Nil(t, err)
		assert.Equal(t, string(ser), v.json)
		var deser ver.DataSource
		err = enc.Decode(ser, &deser)
		assert.Nil(t, err)
	}
}
