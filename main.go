package main

import (
	"fmt"
	"github/MrJharts/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(654)
	fmt.Println(estado)
	fmt.Println(texto)
}