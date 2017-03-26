package tokenizers

type Token struct {
	Position int
	Term     string
}

type Tokenizer func(input string) []*Token
