package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	res := 1
	def := []int{}
	return func() int {
		if len(def) < 2 {
			def = append(def, res)
			return res
		}
		res = res + def[len(def)-2]
		def = append(def, res)
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
