package filters

import (
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type LowercaseFilterSuite struct{}

var _ = check.Suite(&LowercaseFilterSuite{})

func (s *LowercaseFilterSuite) TestLowercaseFilter(c *check.C) {
	expected := []string{"ball"}
	result := LowercaseFilter("Ball")
	c.Check(result, check.DeepEquals, expected)
}
