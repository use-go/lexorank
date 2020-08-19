package lexorank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessEmptyPrevEmptyNext(t *testing.T) {
	rank, ok := Rank("0:47db1", "0:56f21")
	t.Log("ranks:", rank)
	assert.Equal(t, true, ok)
}
