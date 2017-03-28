package indexing

import (
	"github.com/dwayhs/go-search-engine/analysis"
	"github.com/dwayhs/go-search-engine/core"
)

// IndexStore structs implement the indexing and storage of documents with given terms
type IndexStore interface {
	Index(terms []analysis.Term, document core.Document) error
	Search(terms []analysis.Term) []core.Document
}
