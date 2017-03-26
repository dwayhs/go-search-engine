package indexing

import (
	"testing"

	"github.com/dwayhs/go-search-engine/analysis"
	"github.com/dwayhs/go-search-engine/core"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type InvertedIndexSuite struct {
	DocA  core.Document
	Terms []analysis.Term
}

var _ = check.Suite(&InvertedIndexSuite{})

func (s *InvertedIndexSuite) SetUpSuite(c *check.C) {
	s.DocA = core.Document{
		UID: 1,
		Attributes: map[string]string{
			"body": "The quick brown fox jumps over the lazy dog",
		},
	}

	s.Terms = []analysis.Term{
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
}

func (s *InvertedIndexSuite) TestIndex(c *check.C) {
	invertedIndex := NewInvertedIndex()

	invertedIndex.Index(s.Terms, s.DocA)

	expectedInvertedIndex := map[string]TermIncidences{
		"the": TermIncidences{
			Incidences: map[uint32]DocumentTermIncidences{
				1: DocumentTermIncidences{Incidences: []int{1, 7}},
			},
		},
		"quick": TermIncidences{
			Incidences: map[uint32]DocumentTermIncidences{
				1: DocumentTermIncidences{Incidences: []int{2}},
			},
		},
		"brown": TermIncidences{
			Incidences: map[uint32]DocumentTermIncidences{
				1: DocumentTermIncidences{Incidences: []int{3}},
			},
		},
		"fox": TermIncidences{
			Incidences: map[uint32]DocumentTermIncidences{
				1: DocumentTermIncidences{Incidences: []int{4}},
			},
		},
		"jumps": TermIncidences{
			Incidences: map[uint32]DocumentTermIncidences{
				1: DocumentTermIncidences{Incidences: []int{5}},
			},
		},
		"over": TermIncidences{
			Incidences: map[uint32]DocumentTermIncidences{
				1: DocumentTermIncidences{Incidences: []int{6}},
			},
		},
		"lazy": TermIncidences{
			Incidences: map[uint32]DocumentTermIncidences{
				1: DocumentTermIncidences{Incidences: []int{8}},
			},
		},
		"dog": TermIncidences{
			Incidences: map[uint32]DocumentTermIncidences{
				1: DocumentTermIncidences{Incidences: []int{9}},
			},
		},
	}

	expectedDocumentStore := map[uint32]core.Document{
		1: s.DocA,
	}

	c.Check(invertedIndex.InvertedIndex, check.DeepEquals, expectedInvertedIndex)
	c.Check(invertedIndex.DocumentStore, check.DeepEquals, expectedDocumentStore)
}

func (s *InvertedIndexSuite) TestSearch(c *check.C) {
	invertedIndex := NewInvertedIndex()

	invertedIndex.Index(s.Terms, s.DocA)

	searchTerms := []analysis.Term{
		analysis.Term{Position: 1, Term: "brown"},
		analysis.Term{Position: 2, Term: "fox"},
	}

	documents := invertedIndex.Search(searchTerms)
	expectedDocuments := []core.Document{s.DocA}

	c.Check(documents, check.DeepEquals, expectedDocuments)
}
