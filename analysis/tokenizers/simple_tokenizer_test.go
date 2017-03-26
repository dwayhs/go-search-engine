package tokenizers

import (
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type SimpleTokenizerSuite struct{}

var _ = check.Suite(&SimpleTokenizerSuite{})

func (s *SimpleTokenizerSuite) TestSimpleTokenizer(c *check.C) {

	expected := []*Token{
		&Token{Position: 1, Term: "The"},
		&Token{Position: 2, Term: "lazy"},
		&Token{Position: 3, Term: "dog"},
	}
	result := SimpleTokenizer("The lazy dog")
	c.Check(result, check.DeepEquals, expected)
}
