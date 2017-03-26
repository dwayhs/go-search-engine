package indexing

import "github.com/dwayhs/go-search-engine/core"

type DocumentStore interface {
	Index(document core.Document) error
	Search(query string) []core.Document
}
