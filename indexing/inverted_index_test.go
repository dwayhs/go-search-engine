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
}

func (s *InvertedIndexSuite) TestIndex(c *check.C) {
	invertedIndex := NewInvertedIndex()

	invertedIndex.Index(s.DocATerms, s.DocAUID)

	expectedInvertedIndex := map[string]TermIncidences{
		"the": {
			Incidences: map[uint32]DocumentTermIncidences{
				1: {Incidences: []int{1, 7}},
			},
		},
		"quick": {
			Incidences: map[uint32]DocumentTermIncidences{
				1: {Incidences: []int{2}},
			},
		},
		"brown": {
			Incidences: map[uint32]DocumentTermIncidences{
				1: {Incidences: []int{3}},
			},
		},
		"fox": {
			Incidences: map[uint32]DocumentTermIncidences{
				1: {Incidences: []int{4}},
			},
		},
		"jumps": {
			Incidences: map[uint32]DocumentTermIncidences{
				1: {Incidences: []int{5}},
			},
		},
		"over": {
			Incidences: map[uint32]DocumentTermIncidences{
				1: {Incidences: []int{6}},
			},
		},
		"lazy": {
			Incidences: map[uint32]DocumentTermIncidences{
				1: {Incidences: []int{8}},
			},
		},
		"dog": {
			Incidences: map[uint32]DocumentTermIncidences{
				1: {Incidences: []int{9}},
			},
		},
	}

	c.Check(invertedIndex.InvertedIndex, check.DeepEquals, expectedInvertedIndex)
}

func (s *InvertedIndexSuite) TestSearch(c *check.C) {
	invertedIndex := NewInvertedIndex()

	invertedIndex.Index(s.DocATerms, s.DocAUID)

	searchTerms := []analysis.Term{
		{Position: 1, Term: "brown"},
		{Position: 2, Term: "fox"},
	}

	documents := invertedIndex.Search(searchTerms)
	expectedDocumentUIDs := []uint32{s.DocAUID}

	c.Check(documents, check.DeepEquals, expectedDocumentUIDs)
}
