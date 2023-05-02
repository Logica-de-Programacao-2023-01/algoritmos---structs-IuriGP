package main

import "fmt"

type Retangulo struct {
	Largura int
	altura  int
}

func media(r Retangulo) {
	fmt.Println(r.Largura)
	fmt.Println(r.altura)
	fmt.Println(r.Largura * r.altura)
}

func main() {
	r := Retangulo{Largura: 5, altura: 10}
	media(r)

	a := Aluno{
		nome:  "eu",
		idade: 18,
		notas: []float64{1, 2, 3, 4, 5},
	}
	fmt.Println(media_n(a))
}

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func media_n(a Aluno) float64 {
	var sum float64
	for _, nota := range a.notas {
		sum += nota
	}
	return sum / float64(len(a.notas))
}