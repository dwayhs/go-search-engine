package tokenizers

import (
	"testing"

	"github.com/dwayhs/go-search-engine/analysis"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type SimpleTokenizerSuite struct{}

var _ = check.Suite(&SimpleTokenizerSuite{})

func (s *SimpleTokenizerSuite) TestSimpleTokenizer(c *check.C) {

	expected := []analysis.Term{
		{Position: 1, Term: "The"},
		{Position: 2, Term: "lazy"},
		{Position: 3, Term: "dog"},
	}
	result := SimpleTokenizer("The lazy dog")
	c.Check(result, check.DeepEquals, expected)
}
