package gosearchengine

import (
	"testing"

	"github.com/dwayhs/go-search-engine/analysis/analyzers"
	"github.com/dwayhs/go-search-engine/core"
	"github.com/dwayhs/go-search-engine/indexing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type IndexSuite struct{}

var _ = check.Suite(&IndexSuite{})

func (s *IndexSuite) TestIndex(c *check.C) {
	index := NewIndex(
		indexing.NewInvertedIndex(),
		Mapping{
			Attributes: map[string]analyzers.Analyzer{
				"body": analyzers.NewSimpleAnalyzer(),
			},
		},
	)

	docA := core.Document{
		Attributes: map[string]string{
			"body": "The quick brown fox jumps over the lazy dog",
		},
	}

	index.Index(docA)

	searchResult := index.Search("body", "quick fox")

	expected := []core.Document{docA}
	c.Check(searchResult, check.DeepEquals, expected)
}
