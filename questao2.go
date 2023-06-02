package main

import "fmt"

// Crie uma struct chamada Pessoa com os campos "nome", "idade" e "endereço". O campo "endereço" deve ser uma outra struct com os campos "rua", "número", "cidade" e "estado". Escreva uma função que receba uma Pessoa como parâmetro e imprima seu endereço completo.
type Endereço struct {
	rua    string
	numero int
	cidade string
	estado string
}
type Pessoa struct {
	nome     string
	idade    int
	endereço Endereço
}

func main() {
	p := Pessoa{nome: "izabelle",
		idade: 18,
		endereço: Endereço{
			rua:    "asa sul",
			numero: 106,
			cidade: "Brasilia",
			estado: "DF"},
	}
	ImprimaPessoa(p)
}
func ImprimaPessoa(p Pessoa) {
	fmt.Println("nome", p.nome)
	fmt.Println("idade", p.idade)
	fmt.Println("endereço", p.endereço.rua)
	fmt.Println("endereco", p.endereço.numero)
	fmt.Println("endereco", p.endereço.cidade)
	fmt.Println("endereço", p.endereço.estado)
}
