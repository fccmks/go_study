package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	for e := range strings.Fields(s) {

	}
}

func main() {
	//fmt.Printf("Words are: %q", strings.Fields("this is a string"))
	words := []string{}
	words = strings.Fields("hello e")
	fmt.Println(words)
}
