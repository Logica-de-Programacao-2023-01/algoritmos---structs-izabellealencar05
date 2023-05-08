package main

import "time"

//Usando as mesmas structs do exercício anterior, escreva uma função que receba um título de música como parâmetro e retorne uma lista das Playlists que possuem esse título.

func playlistsComMusica(titulo string, playlists []Playlist) []Playlist {
	var playlistsComMusica []Playlist

	for _, p := range playlists {
		for _, m := range p.Musicas {
			if m.Titulo == titulo {
				playlistsComMusica = append(playlistsComMusica, p)
				break
			}
		}
	}

	return playlistsComMusica
}
func main ()
playlists := []Playlist{
Playlist{
Nome: "Playlist 1",
Musicas: []Musica{
Musica{
Titulo:   "Musica 1",
Artista:  "Artista 1",
Duracao:  3 * time.Minute + 30 * time.Second,
},
Musica{
Titulo:   "Musica 2",
Artista:  "Artista 2",
Duracao:  4 * time.Minute + 15 * time.Second,
		},
	},
},
Playlist{
Nome: "Playlist 2",
Musicas: []Musica{
Musica{
Titulo:   "Musica 3",
Artista:  "Artista 3",
Duracao:  2 * time.Minute + 45 * time.Second,
	},
Musica{
Titulo:   "Musica 4",
Artista:  "Artista 4",
Duracao:  5 * time.Minute + 10 * time.Second,
			},
		},
	},
}

tituloBusca := "Musica 3"
playlistsComMusica := playlistsComMusica(tituloBusca, playlists)
fmt.Printf("Playlists com a música '%s': %v\n", tituloBusca, playlistsComMusica)
