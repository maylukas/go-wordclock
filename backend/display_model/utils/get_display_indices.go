package utils

import (
	"strings"
	"unicode/utf8"
)

//GetDisplayIndices returns the charact indices of a given substring in a charact sequence
//cache: It will cache the result in a provided map to prevent repeated lookups
//searchLast: It can either search the text starting from the start or from the end
func GetDisplayIndices(cache map[string][]int, substr string, characts string, searchLast bool) []int {
	result := cache[substr]
	if result == nil {
		wordLen := utf8.RuneCountInString(substr)
		indices := make([]int, wordLen)
		startIndex := -1
		if searchLast {
			startIndex = strings.LastIndex(characts, substr)
		} else {
			startIndex = strings.Index(characts, substr)
		}

		startIndexRune := utf8.RuneCountInString(string([]byte(characts)[:startIndex]))

		for i := 0; i < wordLen; i++ {
			indices[i] = startIndexRune + i
		}
		cache[substr] = indices
		result = indices
	}
	return result
}
