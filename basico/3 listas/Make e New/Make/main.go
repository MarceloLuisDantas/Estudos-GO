// COMO MAKE FUNCIONA E O QUE ELE FAZ

package main

import "fmt"

func main() {
	// Make gera um slice de tamanho X de uma lista de tamanho Y
	var lista = make([]int, 5, 10) // X = 5, Y = 10
	for i, v := range lista {      // [0, 0, 0, 0, 0]
		fmt.Println(i, v)
		lista[i] = i + 1
	}
	// Sempre que um novo indice for adicionado ao Slice, o seu tamanho muda.
	// Passando de X para X + x, sendo x o nomuero de valores adicionados.
	// Caso o tamanho do Slice ultrapasse o tamanho do Array, é criado um novo Array
	lista = append(lista, 6)
	println()
	for i, v := range lista {
		fmt.Println(i, v)
	}
}
