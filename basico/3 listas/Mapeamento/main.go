package main

import "fmt"

// Função para fazer mapemento simples de uma lista de inteiros
func my_map(f func(x int) int, lista []int) []int {
	var new []int
	for _, v := range lista {
		new = append(new, f(v))
	}
	return new
}

// Função para fazer a filtragem simples de uma lsita de inteiros
func filter(f func(x int) bool, lista []int) []int {
	var new []int
	for _, v := range lista {
		if f(v) {
			new = append(new, v)
		}
	}
	return new
}

// Função para reduzir uma lista a um valor unico do mesmo tipo
func reduce(f func(r, x int) int, lista []int) int {
	var resultado int
	for _, v := range lista {
		resultado = f(resultado, v)
	}
	return resultado
}

func main() {

	var lista = []int{1, 2, 3, 4, 5, 12, 7, 8, 9, 10}

	// Função para MAP
	var soma = func(x int) int {
		return x + 2
	}
	// Função para REDUCE
	var par = func(x int) bool {
		return x%2 == 0
	}
	// Função para REDUCE
	var maior = func(r, x int) int {
		if r >= x {
			return r
		}
		return x
	}

	fmt.Println("Lista mapeada")
	var n_lista = my_map(soma, lista)
	fmt.Println(n_lista)

	fmt.Println("\nLista filtrada")
	var n_lista_filtrada = filter(par, lista)
	fmt.Println(n_lista_filtrada)

	fmt.Println("\nValor reduzido? eu acho")
	var valor_reduzido = reduce(maior, lista)
	fmt.Println(valor_reduzido)

}
