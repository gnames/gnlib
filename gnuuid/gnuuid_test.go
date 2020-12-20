package gnuuid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/gnames/gnlib/gnuuid"
)

func TestGNDomain(t *testing.T) {
	assert.Equal(t, GNDomain.String(), "90181196-fecf-5082-a4c1-411d4f314cda")
}

func TestNil(t *testing.T) {
	assert.Equal(t, Nil.String(), "00000000-0000-0000-0000-000000000000")
}

func TestNew(t *testing.T) {
	assert.Equal(t, New("Homo sapiens").String(),
		"16f235a0-e4a3-529c-9b83-bd15fe722110")
}
