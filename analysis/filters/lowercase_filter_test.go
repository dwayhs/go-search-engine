package filters

import (
	"testing"

	"github.com/dwayhs/go-search-engine/analysis"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type LowercaseFilterSuite struct{}

var _ = check.Suite(&LowercaseFilterSuite{})

func (s *LowercaseFilterSuite) TestLowercaseFilter(c *check.C) {
	expected := []analysis.Term{
		{Position: 1, Term: "ball"},
	}
	result := LowercaseFilter(analysis.Term{Position: 1, Term: "Ball"})
	c.Check(result, check.DeepEquals, expected)
}
