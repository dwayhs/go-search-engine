package indexing

import (
	"github.com/dwayhs/go-search-engine/analysis"
	"github.com/dwayhs/go-search-engine/core"
)

type TermIncidences struct {
	Incidences map[uint32]DocumentTermIncidences
}

type DocumentTermIncidences struct {
	Incidences []int
}

type InvertedIndex struct {
	InvertedIndex map[string]TermIncidences
	DocumentStore map[uint32]core.Document
}

func NewInvertedIndex() *InvertedIndex {
	return &InvertedIndex{
		InvertedIndex: map[string]TermIncidences{},
		DocumentStore: map[uint32]core.Document{},
	}
}

func (i *InvertedIndex) Index(terms []analysis.Term, document core.Document) error {
	err := i.storeDocument(document)
	if err != nil {
		return err
	}

	err = i.addTermsToIndex(terms, document)
	if err != nil {
		return err
	}

	return nil
}

func (i *InvertedIndex) Search(query string) []core.Document {
	return []core.Document{}
}

func (i *InvertedIndex) storeDocument(document core.Document) error {
	i.DocumentStore[document.UID] = document

	return nil
}

func (i *InvertedIndex) addTermsToIndex(terms []analysis.Term, document core.Document) error {
	for _, term := range terms {
		i.addTermToIndex(term, document)
	}

	return nil
}

func (i *InvertedIndex) addTermToIndex(term analysis.Term, document core.Document) error {

	if _, ok := i.InvertedIndex[term.Term]; !ok {
		i.InvertedIndex[term.Term] = TermIncidences{
			Incidences: map[uint32]DocumentTermIncidences{},
		}
	}

	if _, ok := i.InvertedIndex[term.Term].Incidences[document.UID]; !ok {
		i.InvertedIndex[term.Term].Incidences[document.UID] = DocumentTermIncidences{
			Incidences: []int{term.Position},
		}
	} else {
		i.InvertedIndex[term.Term].Incidences[document.UID] = DocumentTermIncidences{
			Incidences: append(
				i.InvertedIndex[term.Term].Incidences[document.UID].Incidences,
				term.Position,
			),
		}
	}

	return nil
}
