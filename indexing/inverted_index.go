package indexing

import (
	"github.com/dwayhs/go-search-engine/analysis"
	"github.com/dwayhs/go-search-engine/core"
)

// TermIncidences values stores the incidentes of terms in documents.
type TermIncidences struct {
	Incidences map[uint32]DocumentTermIncidences
}

// DocumentTermIncidences values stores the incidentes of terms in a specific documents.
type DocumentTermIncidences struct {
	Incidences []int
}

// InvertedIndex values control the inverted incex structure and its document store.
type InvertedIndex struct {
	InvertedIndex map[string]TermIncidences
	DocumentStore map[uint32]core.Document
}

// NewInvertedIndex initializes an InvertedIndex with an empty inverted index and document store.
func NewInvertedIndex() *InvertedIndex {
	return &InvertedIndex{
		InvertedIndex: map[string]TermIncidences{},
		DocumentStore: map[uint32]core.Document{},
	}
}

// Index stores and indexes a document for the given terms.
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

// Search queries the inverted index for documents satisfying the given query.
func (i *InvertedIndex) Search(terms []analysis.Term) []core.Document {
	documentUIDS := i.query(terms)
	return i.fetchDocuments(documentUIDS)
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

func (i *InvertedIndex) query(terms []analysis.Term) []uint32 {
	resultingDocumentsHash := map[uint32]bool{}
	resultingDocuments := make([]uint32, 0, 5)

	for _, term := range terms {
		if termIncidences, ok := i.InvertedIndex[term.Term]; ok {
			for documentUID := range termIncidences.Incidences {
				resultingDocumentsHash[documentUID] = true
			}
		}
	}

	for documentUID := range resultingDocumentsHash {
		resultingDocuments = append(resultingDocuments, documentUID)
	}

	return resultingDocuments
}

func (i *InvertedIndex) fetchDocuments(documentUIDS []uint32) []core.Document {
	resultingDocuments := make([]core.Document, 0, len(documentUIDS))

	for _, documentUID := range documentUIDS {
		resultingDocuments = append(
			resultingDocuments,
			i.DocumentStore[documentUID],
		)
	}

	return resultingDocuments
}
