package gosearchengine

import (
	"testing"

	"github.com/dwayhs/go-search-engine/analysis/analyzers"
	"github.com/dwayhs/go-search-engine/core"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type IndexSuite struct {
	DocA  core.Document
	Index *Index
}

var _ = check.Suite(&IndexSuite{})

func (s *IndexSuite) SetUpSuite(c *check.C) {
	s.Index = NewIndex(
		Mapping{
			Attributes: map[string]analyzers.Analyzer{
				"title": analyzers.NewSimpleAnalyzer(),
				"body":  analyzers.NewSimpleAnalyzer(),
			},
		},
	)

	s.DocA = core.Document{
		UID: 1,
		Attributes: map[string]string{
			"title": "some title",
			"body":  "The quick brown fox jumps over the lazy dog",
		},
	}

	s.Index.Index(s.DocA)
}

func (s *IndexSuite) TestQueryPresentBodyShouldHit(c *check.C) {
	s.Index.Index(s.DocA)

	searchResult := s.Index.Search("body", "quick fox")
	expected := []core.Document{s.DocA}
	c.Check(searchResult, check.DeepEquals, expected)
}

func (s *IndexSuite) TestQueryTitleTermInBodyShouldMiss(c *check.C) {
	s.Index.Index(s.DocA)

	searchResult := s.Index.Search("body", "some")
	expected := []core.Document{}
	c.Check(searchResult, check.DeepEquals, expected)
}

func (s *IndexSuite) TestQueryPresentTitleShouldHit(c *check.C) {
	searchResult := s.Index.Search("title", "some")
	expected := []core.Document{s.DocA}
	c.Check(searchResult, check.DeepEquals, expected)
}

func (s *IndexSuite) TestQueryMissingTitleShouldMiss(c *check.C) {
	searchResult := s.Index.Search("title", "missing")
	expected := []core.Document{}
	c.Check(searchResult, check.DeepEquals, expected)
}
