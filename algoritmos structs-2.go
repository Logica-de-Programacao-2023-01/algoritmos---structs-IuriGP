package main

import "fmt"

type Pessoa struct {
	nome     string
	idade    int
	endereço string
}
type Endereco struct {
	rua    int
	número int
	cidade string
	estado string
}

func endereço_retorno(e Endereco) {
	fmt.Println(e.rua)
	fmt.Println(e.número)
	fmt.Println(e.cidade)
	fmt.Println(e.estado)
}

func main() {
	p := Pessoa{
		nome:  "",
		idade: 12,
		endereço: Endereco{
			rua:    505,
			número: 96,
			cidade: "br",
			estado: "br",
		},
	}
	endereço_retorno(p)
}
