package lexorank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaseJiraString(t *testing.T) {
	pos, ok := ParseJira("0|47db1:123")
	t.Log("ranks:", pos.String())
	assert.Equal(t, true, ok)
}

func TestGenerateaMidRank(t *testing.T) {
	rank, ok := Rank("0|47db1:11", "0|56f21:22")
	t.Log("ranks:", rank)
	assert.Equal(t, true, ok)
}

func TestGenerateaBatchRank(t *testing.T) {
	strArray, ok := RanksMore(10, "0|47db1:11", "0|56f21:22")
	t.Log("ranks more:", strArray)
	assert.Equal(t, true, ok)
}
