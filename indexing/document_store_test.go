package indexing

import (
	"testing"

	"github.com/dwayhs/go-search-engine/core"

	"gopkg.in/check.v1"
)

func TestDocumentStore(t *testing.T) {
	check.TestingT(t)
}

type DocumentStoreSuite struct {
	DocA core.Document
	DocB core.Document
}

var _ = check.Suite(&DocumentStoreSuite{})

func (s *DocumentStoreSuite) SetUpSuite(c *check.C) {
	s.DocA = core.Document{
		UID: 1,
		Attributes: map[string]string{
			"body": "The quick brown fox jumps over the lazy dog",
		},
	}
	s.DocB = core.Document{
		UID: 2,
		Attributes: map[string]string{
			"body": "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		},
	}
}

func (s *DocumentStoreSuite) TestFetch(c *check.C) {
	documentStore := NewDocumentStore()

	documentStore.Store(s.DocA.UID, s.DocA)
	fetchedDocument := documentStore.Fetch(s.DocA.UID)

	c.Check(fetchedDocument, check.DeepEquals, s.DocA)
}

func (s *DocumentStoreSuite) TestFetchDocuments(c *check.C) {
	documentStore := NewDocumentStore()

	documentStore.Store(s.DocA.UID, s.DocA)
	documentStore.Store(s.DocB.UID, s.DocB)

	fetchedDocuments := documentStore.FetchDocuments([]core.DocumentUID{
		s.DocA.UID,
		s.DocB.UID,
	})

	c.Check(fetchedDocuments, check.DeepEquals, []core.Document{s.DocA, s.DocB})
}
