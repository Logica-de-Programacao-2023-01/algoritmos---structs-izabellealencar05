package main

// Crie uma struct chamada Playlist com os campos "nome" e "músicas". O campo "músicas" deve ser um slice de structs, cada uma representando uma música com os campos "título", "artista" e "duração". Escreva uma função que receba uma Playlist como parâmetro e imprima o nome da playlist, o tempo total da playlist e as informações de cada música.
import (
	"fmt"
	"time"
)

type Musica struct {
	titulo  string
	artista string
	duracao time.Duration
}

type Playlist struct {
	nome   string
	musica Musica
}

func main() {
	playlist := Playlist{
		nome: "minha playlist",
		musica: Musica{
			titulo:  "End of youth",
			artista: "ed sheeran",
			duracao: 3*time.Minute + 45*time.Second,
		},
	}
	imprimirPlaylist(playlist)
}

func imprimirPlaylist(p Playlist) {
	fmt.Printf("Playlist: %s\n", p.nome)

	var duracaoTotal time.Duration
	for _, m := range Musica {
		duracaoTotal += m.duracao
		fmt.Printf("- %s (%s) - %s\n", m.titulo, m.artista, m.duracao)
	}
	fmt.Printf("duracao total: %s\n", duracaoTotal)
}
