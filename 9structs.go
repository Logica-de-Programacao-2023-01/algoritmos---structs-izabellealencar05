package main

//Crie uma struct chamada Aluno com os campos "nome", "idade" e "notas". O campo "notas" deve ser um slice de float64, representando as notas que o aluno tirou em uma determinada disciplina. Escreva funções que permitam adicionar ou remover notas do aluno, calcular a média das notas e imprimir o nome, idade e média do aluno.
import "fmt"

type Aluno struct {
	Nome  string
	Idade int
	Notas []float64
}

func (a *Aluno) AdicionarNota(nota float64) {
	a.Notas = append(a.Notas, nota)
}

func (a *Aluno) RemoverNota(index int) {
	if index >= 0 && index < len(a.Notas) {
		a.Notas = append(a.Notas[:index], a.Notas[index+1:]...)
	}
}

func (a Aluno) MediaNotas() float64 {
	if len(a.Notas) == 0 {
		return 0.0
	}

	soma := 0.0
	for _, nota := range a.Notas {
		soma += nota
	}
	return soma / float64(len(a.Notas))
}

func (a Aluno) ImprimirInformacoes() {
	fmt.Printf("Nome: %s\n", a.Nome)
	fmt.Printf("Idade: %d\n", a.Idade)
	fmt.Printf("Média das notas: %.2f\n", a.MediaNotas())
}

func main() {
	a := Aluno{Nome: "João", Idade: 20, Notas: []float64{7.5, 8.0, 6.5}}
	a.AdicionarNota(9.0)
	a.RemoverNota(1)
	a.ImprimirInformacoes()
}
