package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Hard
// Создайте программу на Go, которая будет использовать карту (map) для подсчета количества вхождений каждого слова в заданном тексте.
// Входной текст можно представить как строку.
//
//	Программа должна разбить текст на слова, удалить знаки препинания и привести все слова к нижнему регистру перед подсчетом.
//	Затем программа должна вывести на экран список уникальных слов вместе с количеством их вхождений в текст.

func hm1_hard(t string) {
	var clean string
	textlower := strings.ToLower(t)
	for _, symbol := range textlower {
		if unicode.IsLetter(symbol) || unicode.IsSpace(symbol) {
			clean = clean + string(symbol)
		}
	}

	strSlice := strings.Fields(clean)
	countWords := make(map[string]int)
	for _, word := range strSlice {
		countWords[word] = countWords[word] + 1
	}

	fmt.Println(countWords)
}
