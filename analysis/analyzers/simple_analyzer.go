package analyzers

import "github.com/dwayhs/go-search-engine/analysis/tokenizers"

func SimpleAnalyzer() *Analyzer {
	return &Analyzer{
		Tokenizer: tokenizers.SimpleTokenizer,
	}
}
