// Package gosearchengine implements a search engine
//
// It consists of several structures to implement
// indexing and search of documents.
package gosearchengine

import (
	"github.com/dwayhs/go-search-engine/analysis"
	"github.com/dwayhs/go-search-engine/analysis/analyzers"
	"github.com/dwayhs/go-search-engine/core"
	"github.com/dwayhs/go-search-engine/indexing"
)

// Mapping values control analyzers of each document field
type Mapping struct {
	Attributes map[string]analyzers.Analyzer
}

// Index values control the index structures.
type Index struct {
	InvertedIndexes map[string]*indexing.InvertedIndex
	DocumentStore   *indexing.DocumentStore
	Mapping         Mapping
}

// NewIndex initializes an InvertedIndex with the given Mapping.
func NewIndex(mapping Mapping) *Index {
	return &Index{
		Mapping:         mapping,
		DocumentStore:   indexing.NewDocumentStore(),
		InvertedIndexes: map[string]*indexing.InvertedIndex{},
	}
}

// Index stores and indexes a document in the index.
func (i *Index) Index(document core.Document) {
	i.DocumentStore.Store(document.UID, document)

	documentTerms := i.extractTermsFromDocument(document)

	for attribute, terms := range documentTerms {
		indexStore := i.getAttributeIndexStore(attribute)

		indexStore.Index(terms, document.UID)
	}
}

// Search queries the index for documents satisfying the given query.
func (i *Index) Search(attribute string, query string) []core.Document {
	analyzer := i.Mapping.Attributes[attribute]
	terms := analyzer.Analyze(query)

	indexStore := i.getAttributeIndexStore(attribute)

	documentUIDs := indexStore.Search(terms)
	return i.DocumentStore.FetchDocuments(documentUIDs)
}

func (i *Index) extractTermsFromDocument(document core.Document) map[string][]analysis.Term {
	terms := map[string][]analysis.Term{}

	for attribute := range i.Mapping.Attributes {
		analyzer := i.Mapping.Attributes[attribute]

		if attributeValue, ok := document.Attributes[attribute]; ok {
			attributeTerms := analyzer.Analyze(attributeValue)

			terms[attribute] = attributeTerms
		}
	}

	return terms
}

func (i *Index) getAttributeIndexStore(attribute string) *indexing.InvertedIndex {
	if indexStore, ok := i.InvertedIndexes[attribute]; ok {
		return indexStore
	}

	i.InvertedIndexes[attribute] = indexing.NewInvertedIndex()
	return i.InvertedIndexes[attribute]
}
