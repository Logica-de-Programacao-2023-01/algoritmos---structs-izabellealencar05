package main

import "fmt"

// Crie uma struct chamada Animal com os campos "nome", "espécie", "idade" e "som". Escreva funções que permitam modificar o som que o animal faz e uma função que imprima as informações do animal e o som que ele faz.
type Animal struct {
	Nome    string
	Especie string
	Idade   int
	Som     string
}

func (a *Animal) AlterarSom(som string) {
	a.Som = som
}

func (a *Animal) ImprimirInformacoes() {
	fmt.Printf("Nome: %s\nEspécie: %s\nIdade: %d anos\nSom: %s\n", a.Nome, a.Especie, a.Idade, a.Som)
}

func main() {
	animal := Animal{
		Nome:    "Rex",
		Especie: "Cachorro",
		Idade:   2,
		Som:     "Au Au",
	}

	animal.AlterarSom("Woof Woof")
	animal.ImprimirInformacoes()
}
