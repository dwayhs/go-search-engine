package filters

import "strings"

func LowercaseFilter(term string) []string {
	return []string{strings.ToLower(term)}
}
