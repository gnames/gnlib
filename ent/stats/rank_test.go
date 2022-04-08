package stats_test

import (
	"testing"

	"github.com/gnames/gnlib/ent/stats"
	"github.com/stretchr/testify/assert"
)

func TestRank(t *testing.T) {
	assert.True(t, stats.Empire > stats.Kingdom)
	assert.True(t, stats.Class > stats.SubClass)
}
