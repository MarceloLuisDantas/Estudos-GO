// O QUE O MAP FAZ E COMO DICIONARIOS FUNCIONAM

package main

import "fmt"

func main() {
	// Definindo um map com valores já setados
	// Chave - String
	// Valor - int
	var mp map[string]int = map[string]int{
		"foo": 1,
		"bar": 2,
	}
	fmt.Println("Map 1:", mp) // map[bar:2, foo:1]

	// Definindo um map igual o anterior porem vazio
	mp2 := make(map[string]int)
	// Adicionaod 2 valores ao map
	mp2["foo"] = 1
	mp2["bar"] = 2
	fmt.Println("Map 2:", mp2) // map[bar:2, foo:1]

	fmt.Println("------------------")

	delete(mp2, "foo") // Removendo o valor foo do segundo map
	// Vall = valor caso o item exista
	// Ok = o item existe ou não
	val, ok := mp2["foo"]
	fmt.Println(val, ok) // N tem nada pq foi deletado

	fmt.Println("------------------")

	// Mapas podem ser utilizados em for normalmente como se fossem listas
	for k, v := range mp {
		fmt.Println("Chave:", k, "Valor:", v)
	}

	fmt.Println("------------------")

	// Não sei pq vc taria isso, porem é póssivel
	mp3 := map[string]map[string]int{
		"foo": {
			"foo1": 1,
			"bar1": 2,
		},
		"bar": {
			"foo2": 3,
			"bar2": 4,
		},
	}
	for _, v := range mp3 {
		for k2, v2 := range v {
			fmt.Println("Chave:", k2, "Valor:", v2)
		}
	}
}
