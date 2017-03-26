package filters

import "github.com/dwayhs/go-search-engine/analysis"

type Filter func(input string) []*analysis.Term
