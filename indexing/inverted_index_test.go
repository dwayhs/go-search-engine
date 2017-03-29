package indexing

import (
	"testing"

	"github.com/dwayhs/go-search-engine/analysis"

	"gopkg.in/check.v1"
)

func TestInvertedIndex(t *testing.T) {
	check.TestingT(t)
}

type InvertedIndexSuite struct {
	DocAUID   uint32
	DocATerms []analysis.Term
}

var _ = check.Suite(&InvertedIndexSuite{})

func (s *InvertedIndexSuite) SetUpSuite(c *check.C) {
	s.DocAUID = 1
	s.DocATerms = []analysis.Term{
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

	invertedIndex.Index(s.DocATerms, s.DocAUID)

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

	c.Check(invertedIndex.InvertedIndex, check.DeepEquals, expectedInvertedIndex)
}

func (s *InvertedIndexSuite) TestSearch(c *check.C) {
	invertedIndex := NewInvertedIndex()

	invertedIndex.Index(s.DocATerms, s.DocAUID)

	searchTerms := []analysis.Term{
		analysis.Term{Position: 1, Term: "brown"},
		analysis.Term{Position: 2, Term: "fox"},
	}

	documents := invertedIndex.Search(searchTerms)
	expectedDocumentUIDs := []uint32{s.DocAUID}

	c.Check(documents, check.DeepEquals, expectedDocumentUIDs)
}
