package filters

import (
	"strings"

	"github.com/dwayhs/go-search-engine/analysis"
)

// LowercaseFilter processes a list of terms and lowercases each term
func LowercaseFilter(term analysis.Term) []analysis.Term {
	return []analysis.Term{
		{Position: term.Position, Term: strings.ToLower(term.Term)},
	}
}
