package filters

import "github.com/dwayhs/go-search-engine/analysis"

type Filter func(term analysis.Term) []analysis.Term
