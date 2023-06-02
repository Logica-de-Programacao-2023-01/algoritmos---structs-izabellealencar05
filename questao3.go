package main

import "fmt"

// Crie uma struct chamada Triângulo com os campos "base" e "altura". Escreva uma função que receba um Triângulo como parâmetro e calcule a área do triângulo (área = base * altura / 2).
type Triangulo struct {
	base   float64
	altura float64
}

func main() {
	p := Triangulo{altura: 6, base: 5}
	fmt.Println(AreaDoTriangulo(p))
}

func AreaDoTriangulo(p Triangulo) (area float64) {
	area = (p.base * p.altura) / 2
	return area

}
