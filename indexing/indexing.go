package indexing

import (
	"github.com/dwayhs/go-search-engine/analysis"
	"github.com/dwayhs/go-search-engine/core"
)

type IndexStore interface {
	Index(terms []analysis.Term, document core.Document) error
	Search(terms []analysis.Term) []core.Document
}
