package tokenizers

type Term struct {
	Position int
	Term     string
}

type Tokenizer func(input string) []*Term
