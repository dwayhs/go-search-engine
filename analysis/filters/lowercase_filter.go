package filters

import (
	"strings"

	"github.com/dwayhs/go-search-engine/analysis"
)

func LowercaseFilter(term analysis.Term) []analysis.Term {
	return []analysis.Term{
		analysis.Term{Position: term.Position, Term: strings.ToLower(term.Term)},
	}
}
