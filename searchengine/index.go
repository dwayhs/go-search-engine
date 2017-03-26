package searchengine

import (
	"github.com/dwayhs/go-search-engine/analysis"
	"github.com/dwayhs/go-search-engine/analysis/analyzers"
	"github.com/dwayhs/go-search-engine/core"
	"github.com/dwayhs/go-search-engine/indexing"
)

type Mapping struct {
	Attributes map[string]analyzers.Analyzer
}

type Index struct {
	IndexStore indexing.IndexStore
	Mapping    Mapping
}

func NewIndex(indexStore indexing.IndexStore, mapping Mapping) *Index {
	return &Index{
		IndexStore: indexStore,
		Mapping:    mapping,
	}
}

func (i *Index) Index(document core.Document) {
	terms := i.extractTermsFromDocument(document)

	i.IndexStore.Index(terms, document)
}

func (i *Index) Search(attribute string, query string) []core.Document {
	analyzer := i.Mapping.Attributes[attribute]
	terms := analyzer.Analyze(query)

	return i.IndexStore.Search(terms)
}

func (i *Index) extractTermsFromDocument(document core.Document) []analysis.Term {
	terms := make([]analysis.Term, 0, 20)

	for attribute := range i.Mapping.Attributes {
		analyzer := i.Mapping.Attributes[attribute]

		if attributeValue, ok := document.Attributes[attribute]; ok {
			attributeTerms := analyzer.Analyze(attributeValue)
			terms = append(terms, attributeTerms...)
		}

	}

	return terms
}
