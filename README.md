# Go Search Engine

[![Build Status](https://travis-ci.org/dwayhs/go-search-engine.svg?branch=master)](https://travis-ci.org/dwayhs/go-search-engine)
[![Go Report Card](https://goreportcard.com/badge/github.com/dwayhs/go-search-engine)](https://goreportcard.com/report/github.com/dwayhs/go-search-engine)
[![BCH compliance](https://bettercodehub.com/edge/badge/dwayhs/go-search-engine?branch=master)](https://bettercodehub.com/)
[![GoDoc](https://godoc.org/github.com/dwayhs/go-search-engine?status.svg)](https://godoc.org/github.com/dwayhs/go-search-engine)

This project is intended for educational purposes.

## Usage

### Creating an index

```go
index := NewIndex(
  Mapping{
    Attributes: map[string]analyzers.Analyzer{
      "body": analyzers.NewSimpleAnalyzer(),
    },
  },
)
```

### Indexing

```go
docA := core.Document{
  Attributes: map[string]string{
    "body": "The quick brown fox jumps over the lazy dog",
  },
}

index.Index(docA)
```

### Querying

```go
searchResult := index.Search("body", "quick fox")
```

# Credits

Project inspired by the presentation ["Building A Python-Based Search Engine"](https://www.youtube.com/watch?v=cY7pE7vX6MU) on Pycon
US 2012 by Daniel Lindsley.
