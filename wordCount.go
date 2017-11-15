package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	words := []string{}
	count := 1
	words = strings.Fields(s)
	for i := 0; i < len(words); i++ {
		_, ok := wordMap[words[i]]
		if ok {
			wordMap[words[i]] = count + 1
		} else {
			wordMap[words[i]] = count
		}
	}
	return wordMap
}

func main() {
	//fmt.Printf("Words are: %q", strings.Fields("this is a string"))
	fmt.Println(WordCount("hi hi go he"))
}
