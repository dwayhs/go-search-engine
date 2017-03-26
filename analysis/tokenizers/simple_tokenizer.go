package tokenizers

import "strings"

func SimpleTokenizer(input string) []*Token {
	terms := strings.Split(input, " ")
	tokens := make([]*Token, 0, 20)

	for position, term := range terms {
		tokens = append(tokens, &Token{
			Position: position + 1,
			Term:     term,
		})
	}

	return tokens
}
