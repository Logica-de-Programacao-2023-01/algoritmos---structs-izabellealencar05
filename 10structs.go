package main

import "fmt"

// Crie uma struct chamada Filme com os campos "título", "diretor", "ano" e "avaliações". O campo "avaliações" deve ser um slice de inteiros representando as notas que o filme recebeu dos espectadores. Escreva funções que permitam adicionar ou remover avaliações do filme, calcular a média das avaliações e imprimir as informações do filme e sua média de avaliações.
type Filme struct {
	Título     string
	Diretor    string
	Ano        int
	Avaliações []int
}

func (f *Filme) AdicionarAvaliação(avaliação int) {
	f.Avaliações = append(f.Avaliações, avaliação)
}

func (f *Filme) RemoverAvaliação(index int) {
	if index >= 0 && index < len(f.Avaliações) {
		f.Avaliações = append(f.Avaliações[:index], f.Avaliações[index+1:]...)
	}
}

func (f Filme) MédiaAvaliações() float64 {
	if len(f.Avaliações) == 0 {
		return 0.0
	}

	soma := 0.0
	for _, avaliação := range f.Avaliações {
		soma += float64(avaliação)
	}
	return soma / float64(len(f.Avaliações))
}

func (f Filme) ImprimirInformações() {
	fmt.Printf("Título: %s\n", f.Título)
	fmt.Printf("Diretor: %s\n", f.Diretor)
	fmt.Printf("Ano: %d\n", f.Ano)
	fmt.Printf("Média das avaliações: %.2f\n", f.MédiaAvaliações())
}

func main() {
	f := Filme{Título: "Matrix", Diretor: "The Wachowski Brothers", Ano: 1999, Avaliações: []int{8, 9, 7, 10}}
	f.AdicionarAvaliação(8)
	f.RemoverAvaliação(1)
	f.ImprimirInformações()
}
