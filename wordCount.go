package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	words := []string{}
	words = strings.Fields(s)
	for i := 0; i < len(words); i++ {
		count, ok := wordMap[words[i]]
		if ok {
			wordMap[words[i]] = count + 1
		} else {
			wordMap[words[i]] = 1
		}
	}
	return wordMap
}

func main() {
	//fmt.Printf("Words are: %q", strings.Fields("this is a string"))
	fmt.Println(WordCount("hi hi hi hi go he go"))
}
