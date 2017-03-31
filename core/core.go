package core

// DocumentUID values represent the id of a document
type DocumentUID uint32

// Document values represent a document that can be indexed and queried
type Document struct {
	UID        DocumentUID
	Attributes map[string]string
}
