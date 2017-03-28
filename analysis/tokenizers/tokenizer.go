package tokenizers

import "github.com/dwayhs/go-search-engine/analysis"

// Tokenizer functions translate an input string into a list of terms
type Tokenizer func(input string) []analysis.Term
