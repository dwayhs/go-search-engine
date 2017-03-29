package indexing

import (
	"github.com/dwayhs/go-search-engine/core"
)

// DocumentStore values control the storage of documents.
type DocumentStore struct {
	StoreMap map[uint32]core.Document
}

// NewDocumentStore initializes a DocumentStore with an empty store.
func NewDocumentStore() *DocumentStore {
	return &DocumentStore{
		StoreMap: map[uint32]core.Document{},
	}
}

// Store stores a document for the given UID.
func (i *DocumentStore) Store(UID uint32, document core.Document) {
	i.StoreMap[UID] = document
}

// Fetch retrieves a document with the given documentUID.
func (i *DocumentStore) Fetch(documentUID uint32) core.Document {
	return i.StoreMap[documentUID]
}

// FetchDocuments retrieves a list of document with the given documentUIDS.
func (i *DocumentStore) FetchDocuments(documentUIDS []uint32) []core.Document {
	resultingDocuments := make([]core.Document, 0, len(documentUIDS))

	for _, documentUID := range documentUIDS {
		resultingDocuments = append(resultingDocuments, i.Fetch(documentUID))
	}

	return resultingDocuments
}
