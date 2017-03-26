package analyzers

import (
	"github.com/dwayhs/go-search-engine/analysis"
	"github.com/dwayhs/go-search-engine/analysis/filters"
	"github.com/dwayhs/go-search-engine/analysis/tokenizers"
)

type Analyzer struct {
	Tokenizer tokenizers.Tokenizer
	Filters   []filters.Filter
}

func (a *Analyzer) Analyze(input string) []*analysis.Term {
	return a.Tokenizer(input)
}
