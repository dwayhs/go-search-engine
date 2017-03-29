package analyzers

import (
	"testing"

	"github.com/dwayhs/go-search-engine/analysis"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type SimpleAnalyzerSuite struct{}

var _ = check.Suite(&SimpleAnalyzerSuite{})

func (s *SimpleAnalyzerSuite) TestLowercaseFilter(c *check.C) {
	expected := []analysis.Term{
		{Position: 1, Term: "the"},
		{Position: 2, Term: "quick"},
		{Position: 3, Term: "brown"},
		{Position: 4, Term: "fox"},
		{Position: 5, Term: "jumps"},
		{Position: 6, Term: "over"},
		{Position: 7, Term: "the"},
		{Position: 8, Term: "lazy"},
		{Position: 9, Term: "dog"},
	}

	analyzer := NewSimpleAnalyzer()
	result := analyzer.Analyze("The quick brown fox jumps over the lazy dog")
	c.Check(result, check.DeepEquals, expected)
}
