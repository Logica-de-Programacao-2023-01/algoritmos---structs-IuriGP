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
	r := Retangulo{Largura:5 ,altura: 10 }
	media(r)
}

type Retangle struct {
	width  int
	height int
}

func () {
	fmt.Println(" ")
	fmt.Scan(&width)
	fmt.Println(" ")
	fmt.Scan(&height)

	fmt.print(area())
}
func area (r Rectangle) int {
	return w.width * w.height
}
