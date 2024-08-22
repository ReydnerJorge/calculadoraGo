package main

import "fmt"

func main() {
	x := soma(2, 4, 6)
	y := subtracao(10, 20)
	w := multiplicacao(100, 100)
	z := divisao(100)

	fmt.Println(x, y, w, z)
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}
func subtracao(i ...int) int {
	total := 0
	for _, v := range i {
		total = v - total
	}
	return total
}
func multiplicacao(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}
func divisao(i ...int) int {
	total := 2
	for _, v := range i {
		total = v / total
	}
	return total
}
