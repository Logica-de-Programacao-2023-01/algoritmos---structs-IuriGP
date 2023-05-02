package main

import "fmt"

type triangulo struct {
	Largura int
	altura  int
}

func area(t triangulo) {
	//fmt.Println(t.Largura)
	//fmt.Println(t.altura)
	fmt.Println(t.Largura * t.altura / 2)
}

func main() {
	t := triangulo{Largura: 5, altura: 10}
	area(t)
}
