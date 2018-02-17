package leddisplay

import (
	"strings"
	"unicode/utf8"
)

func getDisplayIndices(cacheArray map[string][]int, word string, characts string, searchLast bool) []int {
	result := cacheArray[word]
	if result == nil {
		wordLen := utf8.RuneCountInString(word)
		indices := make([]int, wordLen)
		startIndex := -1
		if searchLast {
			startIndex = strings.LastIndex(characts, word)
		} else {
			startIndex = strings.Index(characts, word)
		}

		startIndexRune := utf8.RuneCountInString(string([]byte(characts)[:startIndex]))

		for i := 0; i < wordLen; i++ {
			indices[i] = startIndexRune + i
		}
		cacheArray[word] = indices
		result = indices
	}
	return result
}
