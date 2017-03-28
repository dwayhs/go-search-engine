package analyzers

import (
	"github.com/dwayhs/go-search-engine/analysis"
	"github.com/dwayhs/go-search-engine/analysis/filters"
	"github.com/dwayhs/go-search-engine/analysis/tokenizers"
)

// Analyzer values control the analyzer process, first tokenize than filter each token
type Analyzer struct {
	Tokenizer tokenizers.Tokenizer
	Filters   []filters.Filter
}

// Analyze executes the analisis process and returns a list of terms
func (a *Analyzer) Analyze(input string) []analysis.Term {
	terms := a.Tokenizer(input)

	for _, filter := range a.Filters {
		terms = applyFilter(filter, terms)
	}

	return terms
}

func applyFilter(filter filters.Filter, terms []analysis.Term) []analysis.Term {
	filteredTerms := make([]analysis.Term, 0, 5)

	for _, term := range terms {
		filteredTerms = append(filteredTerms, filter(term)...)
	}

	return filteredTerms
}
