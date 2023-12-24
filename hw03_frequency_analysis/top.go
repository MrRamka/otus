package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const AmountWords = 10

func Top10(inputText string) []string {
	frequencyMap := make(map[string]int)

	words := strings.Fields(inputText)

	for _, word := range words {
		frequencyMap[word]++
	}

	uniqueWords := make([]string, 0)

	for k := range frequencyMap {
		uniqueWords = append(uniqueWords, k)
	}

	sort.Slice(uniqueWords, func(i, j int) bool {
		firstElement := uniqueWords[i]
		firstElementAmount := frequencyMap[firstElement]
		secondElement := uniqueWords[j]
		secondElementAmount := frequencyMap[secondElement]

		// Если количество одинаковое, то сортируем лексикографически
		if firstElementAmount == secondElementAmount {
			return firstElement < secondElement
		}
		return firstElementAmount > secondElementAmount
	})

	return uniqueWords[:AmountWords]
}
