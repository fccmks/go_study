package main

import "fmt"

//первое число - индекс, второе не число а слайс
func Poc(x, y int) [][]uint8 {
	spic := [][]uint8{}
	for i := 0; i < y; i++ {
		ap := []uint8{}
		for j := 0; j < x; j++ {
			ap = append(ap, uint8(j*x))
		}
		spic = append(spic, ap)
	}
	return spic
}

func main() {
	for re, v := range Poc(3, 3) {
		fmt.Println(re, v)
	}
}
