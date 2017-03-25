package tokenizers

import "strings"

func SimpleTokenizer(term string) []string {
	return strings.Split(term, " ")
}
