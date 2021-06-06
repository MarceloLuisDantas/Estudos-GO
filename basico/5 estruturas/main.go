package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Estrutura pessoa com nome e idade
type Pessoa struct {
	idade int8
	nome  string
}

// Função que retorna uma estrutura
func new_pessoa(idade int8, nome string) Pessoa {
	// Criando uma variavel com uma estrutura
	var pessoa Pessoa = Pessoa{
		idade,
		nome,
	}
	return pessoa
}

// Função de input cru
func input(label string, s *bufio.Scanner) string {
	fmt.Printf(label)
	s.Scan()
	return s.Text()
}

// Função para faciliatar entrade de valores
func int_input(label string, s *bufio.Scanner) (int64, error) {
	// Pega a String e covnerte em numero
	return strconv.ParseInt(input(label, s), 10, 8)
}

// Função que faz input de uma String sem espaços adicionais
func str_input(label string, s *bufio.Scanner) string {
	// Pega a String da entrada e splita ela
	entrada := strings.Split(input(label, s), " ")
	var final []string
	// Percore a lista splitada e retorna apenas as palavras
	for _, v := range entrada {
		if v != "" {
			final = append(final, v)
		}
	}
	// Junta a lista em uma unica String
	return strings.Join(final, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	nome := str_input("Digite seu nome: ", scanner)
	idade, err := int_input("Digite sua idade: ", scanner)
	if err != nil {
		panic("Entrada invalida")
	}

	p := Pessoa{int8(idade), nome}
	fmt.Println(p)
}
