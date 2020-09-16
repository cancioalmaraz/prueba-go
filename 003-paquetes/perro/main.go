package main

import (
	"fmt"

	"github.com/cancioalmaraz/prueba-go/003-paquetes/gato"
)

func main() {
	fmt.Println("Hola desde perro")

	gato.Hola()
	gato.Comen()
}
