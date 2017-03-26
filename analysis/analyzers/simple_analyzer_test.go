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
		analysis.Term{Position: 1, Term: "the"},
		analysis.Term{Position: 2, Term: "quick"},
		analysis.Term{Position: 3, Term: "brown"},
		analysis.Term{Position: 4, Term: "fox"},
		analysis.Term{Position: 5, Term: "jumps"},
		analysis.Term{Position: 6, Term: "over"},
		analysis.Term{Position: 7, Term: "the"},
		analysis.Term{Position: 8, Term: "lazy"},
		analysis.Term{Position: 9, Term: "dog"},
	}

	analyzer := SimpleAnalyzer()
	result := analyzer.Analyze("The quick brown fox jumps over the lazy dog")
	c.Check(result, check.DeepEquals, expected)
}
