package main

import "fmt"

// Crie uma struct chamada Viagem com os campos "origem", "destino", "data" e "preço". Escreva uma função que receba um slice de Viagens como parâmetro e retorne a viagem mais cara.
type Viagem struct {
	Origem  string
	Destino string
	Data    string
	Preco   float64
}

func viagemMaisCara(viagens []Viagem) Viagem {
	if len(viagens) == 0 {
		return Viagem{}
	}

	viagemMaisCara := viagens[0]
	for _, viagem := range viagens {
		if viagem.Preco > viagemMaisCara.Preco {
			viagemMaisCara = viagem
		}
	}

	return viagemMaisCara
}

func main() {
	viagens := []Viagem{
		{Origem: "São Paulo", Destino: "Nova York", Data: "01/06/2023", Preco: 5000.0},
		{Origem: "Rio de Janeiro", Destino: "Miami", Data: "15/06/2023", Preco: 6000.0},
		{Origem: "Belo Horizonte", Destino: "Paris", Data: "20/06/2023", Preco: 7000.0},
		{Origem: "Salvador", Destino: "Londres", Data: "30/06/2023", Preco: 8000.0},
	}

	viagemMaisCara := viagemMaisCara(viagens)
	fmt.Printf("A viagem mais cara é de %s para %s na data %s, com preço de R$ %.2f\n", viagemMaisCara.Origem, viagemMaisCara.Destino, viagemMaisCara.Data, viagemMaisCara.Preco)
}
