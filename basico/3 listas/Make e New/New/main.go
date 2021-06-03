// COMO O NEW FUNCIONA E O QUE ELE FAZ

package main

import "fmt"

// Recebe um ponteiro de um inteiro
func teste(num *int) {
	// Num contei o endereço que de memoria do valor criado
	// *num mostra o valor atrelado a esse ponteiro
	fmt.Println("Endereço:", num, "Valor:", *num)
}

// Recebe um inteiro normal
func teste2(num int) {
	// &num mostra o endereço de memoria a onde num esta alocado
	fmt.Println("Endereço:", &num, "Valor:", num)
}

func main() {
	// Cria um ponteiro para um valor int vazio(0)
	num := new(int)
	fmt.Println("Endereço:", num, "Valor:", *num) // Endereço X : 0
	teste(num)                                    // Endereço X : 0
	teste2(*num)                                  // Endereço Y : 0
	// Os valores são os mesmos, porem endereço de memoria em
	// teste e teste2 são diferentes, pois em teste o que
	// passamos é o ponteiro, e não uma copia, logo é mais rapido
	// porem no teste 2 nos estamos passando o valor atrelado ao
	// ponteiro, logo fazemos uma copia, o que torna mais facil
	// manuzear porem torna mais lento

}
