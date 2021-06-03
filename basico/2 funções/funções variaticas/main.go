package main

import "fmt"

func soma(num ...int) int {
	var t int
	for _, v := range num {
		t += v
	}
	return t
}

func main() {
	var lista = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(soma(lista...))

	fmt.Println(soma(1, 2, 3, 4, 5, 6, 7, 8))

}
