package analyzers

import (
	"github.com/dwayhs/go-search-engine/analysis/filters"
	"github.com/dwayhs/go-search-engine/analysis/tokenizers"
)

// NewSimpleAnalyzer creates an analyzer that uses SimpleTokenizer and LowercaseFilter
func NewSimpleAnalyzer() Analyzer {
	return Analyzer{
		Tokenizer: tokenizers.SimpleTokenizer,
		Filters: []filters.Filter{
			filters.LowercaseFilter,
		},
	}
}
