// SLICES DE ARRAYS EM GO

package main

import "fmt"

// O tipo Slice trabalha com referencias, um slice de um array não
// é uma copia dos valores, e sim uma referencia. Logo se alterarmos
// o array original, o Slice que esta apontando para ele sera alterado
type SliceHeader struct {
	Deta *int
	Len  int
}

func main() {
	// Array de inteiro
	var ar_1 = [3]int{1, 2, 3} // [1, 2, 3]
	var ar_2 [3]int            // [0, 0, 0]
	// Arrays precisa ser declarado já sabendo o tamanho que eles
	// deveram ter durante a execução do codigo, são seguros
	// porem não maleaveis
	fmt.Println("\nLista 1: ")
	mostra_lista(ar_1)
	fmt.Println("\nLista 2: ")
	mostra_lista(ar_2)

	// Slice de inteiro
	var sl_1 = []int{1, 2, 3} // [1, 2, 3]
	var sl_2 []int            // []
	// Slices não precisam ter umtamanho predeterminado, são mais
	// maleaveis porem são menos seguros e mais lentos
	fmt.Println("\nSlice 1: ")
	mostra_slice(sl_1)
	fmt.Println("\nSlice 2: ")
	mostra_slice(sl_2)

	println()
	// Criando uma lista
	var idades = [5]int{4, 12, 18, 5, 47}
	// Criando um slice baseado na lista anterior
	var idades_s = idades[0:5] // Todos os valores "copiados"
	for i := range idades {
		fmt.Println("Valor   ", i, "lista =", idades[i], "| Slice =", idades_s[i])
		fmt.Println("Ponteiro", i, &idades[i], &idades_s[i])
		fmt.Println("----------------------------------------")
	}

	println()

	fmt.Println("Lista: ", idades)
	fmt.Println("Slice: ", idades_s)
	idades[1] = 0

	// Ao alterar o valor 1 na lista original o valor do slice
	// é alterado, pois como dito, não é uma copia, é uma referencia
	fmt.Println("2 valor da lista alterada")
	fmt.Println("Lista: ", idades)
	fmt.Println("Slice: ", idades_s)
}

// Caso se passe um slice para uma função, mesmo que n fique claro
// o parametro n sera um clone, logo alterar uma slice dentro do
// escopo da função ira alterar o a lista original
func mostra_slice(x []int) {
	if len(x) == 0 {
		fmt.Println("Slice vazio")
	} else {
		for i, v := range x {
			fmt.Println("Pos", i+1, "=", v)
		}
	}
}

func mostra_lista(x [3]int) {
	for i, v := range x {
		fmt.Println("Pos", i+1, "=", v)
	}
}

func mostra_ponteiros(idade [5]int, slice []int) {

}
