package lexorank

import (
	"fmt"
	"regexp"
)

const (
	orderToByte  = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	maxMultiRank = 61 //10 + 26 + 26 - 1

	trailer = "UUUUUUUU"

	minChar = byte('0')
	maxChar = byte('z')
)

////    <bucket>|<base36>[:<base36>]
var jiraRank = regexp.MustCompile(`^([012])\|([0-9a-z]+)(:[0-9a-z]*)?$`)

//Position Parse a Jira (Cloud?) lexorank field, which seems to look like:
//    <bucket>|<base36>[:<base36>]
//Position for a position
type Position struct {
	Bucket byte
	Major  string
	Minor  string // note this includes the ":" prefix
}

//String for result
func (p Position) String() string {
	return fmt.Sprintf("%d|%s%s", p.Bucket, p.Major, p.Minor)
}
