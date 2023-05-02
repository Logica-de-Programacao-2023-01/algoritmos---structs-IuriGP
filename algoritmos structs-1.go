package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func calculo(r Circulo) {
	fmt.Println(math.Pi * (r.raio * r.raio))
}

func main() {
	p := Circulo{raio: 5}
	calculo(p)
}
