package main

import "fmt"

type Personagem struct {
	Nome string
	Vida int
}

type Mago struct {
	Player Personagem
}

type Necromante struct {
	Player Personagem
}

type Suporte struct {
	Player Personagem
}

func main() {
	mirra := Mago{
		Personagem{
			"Mirra",
			100000,
		},
	}
	comprar_posao(mirra)
	comprar_po_morto(mirra)
}

func comprar_posao(s ISuporte) {
	fmt.Println("Todos os magos e clerigos deste reino são bem vindos em minha loja.")
	fmt.Println("Compre o que bem entender e precisar.")
}

func comprar_po_morto(i IEnvacador) {
	fmt.Println("Apenas os mais fortes magos e lichs podem usar dos meus itens.")
	fmt.Println("Cuidado, ate mesmo a imortalidade pode ter um fim.")
}

func (m Mago) summonar() {
	fmt.Println("EU INVOCO, OS DRAGÕES DE NORKIENN")
}
func (n Necromante) summonar() {
	fmt.Println("EU INVOCO, OS MORTOS DE JARZIKI")
}
func (m Mago) curar(p *Personagem) {
	p.Vida += 10
	fmt.Printf("Eu curo %s em 10\n", p.Nome)
}
func (m Suporte) curar(p *Personagem) {
	p.Vida += 100
	fmt.Printf("Eu curo %s completamente\n", p.Nome)
}

type IEnvacador interface {
	summonar()
}

type ISuporte interface {
	curar(p *Personagem)
}
