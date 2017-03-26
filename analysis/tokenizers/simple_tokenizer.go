package tokenizers

import "strings"

func SimpleTokenizer(input string) []*Term {
	terms := strings.Split(input, " ")
	tokens := make([]*Term, 0, 20)

	for position, term := range terms {
		tokens = append(tokens, &Term{
			Position: position + 1,
			Term:     term,
		})
	}

	return tokens
}
