package lexorank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessEmptyPrevEmptyNext(t *testing.T) {
	rank, ok := Rank("", "")
	assert.Equal(t, "U", rank)
	assert.Equal(t, true, ok)
}

func TestSuccessEmptyPrev(t *testing.T) {

}

func TestSuccessEmptyNext(t *testing.T) {

}

func TestSuccessNewDigit(t *testing.T) {

}

func TestSuccessMidValue(t *testing.T) {

}

func TestSuccessNewDigitMidValue(t *testing.T) {

}

func TestFailSamePrevNext(t *testing.T) {

}

func TestFailAdjacent(t *testing.T) {

}
