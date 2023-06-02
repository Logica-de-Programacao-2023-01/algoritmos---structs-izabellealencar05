package main

import (
	"fmt"
	"time"
)

//Crie uma struct chamada Funcionário com os campos "nome", "salário" e "idade". Escreva funções que permitam aumentar ou diminuir o salário do funcionário em uma determinada porcentagem e uma função que calcule o tempo de serviço do funcionário (considerando que ele começou a trabalhar aos 18 anos).
type Funcionario struct {
	Nome    string
	Salario float64
	Idade   int
}

func (f *Funcionario) AumentarSalario(porcentagem float64) {
	aumento := f.Salario * (porcentagem / 100)
	f.Salario += aumento
}

func (f *Funcionario) DiminuirSalario(porcentagem float64) {
	desconto := f.Salario * (porcentagem / 100)
	f.Salario -= desconto
}

func (f *Funcionario) TempoServico() int {
	anosTrabalhados := time.Now().Year() - (f.Idade + 18)
	return anosTrabalhados
}
}
func main {

	funcionario := Funcionario{
		Nome:    "João",
		Salario: 3000.0,
		Idade:   25,
	}

	fmt.Println("Salário atual:", funcionario.Salario)
	funcionario.AumentarSalario(10.0)
	fmt.Println("Salário após aumento de 10%:", funcionario.Salario)
	funcionario.DiminuirSalario(5.0)
	fmt.Println("Salário após desconto de 5%:", funcionario.Salario)
	tempoServico := funcionario.TempoServico()
	fmt.Println("Tempo de serviço:", tempoServico, "anos")
}