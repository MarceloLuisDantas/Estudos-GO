package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int8
}

// Metodos usando atributos por copia
func (p Pessoa) beber() bool {
	return p.Idade >= 18
}

// Metodos usando atributos por referencia
func (p *Pessoa) aniversario() {
	p.Idade += 1
	fmt.Println("Feliz aniversario", p.Nome)
	fmt.Println("Agora você tem", p.Idade, "anos")
}

func main() {
	var p Pessoa = Pessoa{"Marcelo Luis", 19}
	if p.beber() {
		fmt.Println("Você já pode beber.")
	} else {
		fmt.Println("Você não pdoe beber.")
	}
	p.aniversario()
}
