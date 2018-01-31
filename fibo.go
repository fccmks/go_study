package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibo(c int8) []int8 {
	res := []int8{0}
	if c == 0 {return res}
	if c == 1 {
		res = append(res, c)
	}
	return res
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

