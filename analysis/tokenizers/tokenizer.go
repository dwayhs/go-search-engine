package tokenizers

import "github.com/dwayhs/go-search-engine/analysis"

type Tokenizer func(input string) []analysis.Term
