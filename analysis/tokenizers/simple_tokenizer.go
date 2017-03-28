package tokenizers

import (
	"strings"

	"github.com/dwayhs/go-search-engine/analysis"
)

// SimpleTokenizer splits tokens by spaces and returns a list of terms
func SimpleTokenizer(input string) []analysis.Term {
	terms := strings.Split(input, " ")
	tokens := make([]analysis.Term, 0, 20)

	for position, term := range terms {
		tokens = append(tokens, analysis.Term{
			Position: position + 1,
			Term:     term,
		})
	}

	return tokens
}
