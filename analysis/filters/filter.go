package filters

import "github.com/dwayhs/go-search-engine/analysis"

// Filter functions processes a list of terms and returns a new list of processed terms
type Filter func(term analysis.Term) []analysis.Term
