package context_test

import (
	"testing"

	"github.com/gnames/gnlib/ent/context"
	"github.com/stretchr/testify/assert"
)

func TestRank(t *testing.T) {
	assert.True(t, context.Family > context.Genus)
	assert.True(t, context.Kingdom > context.Phylum)
}