package core

// Document values represent a document that can be indexed and queried
type Document struct {
	UID        uint32
	Attributes map[string]string
}
