package main

import "fmt"

type persona struct {
	Nombre, Apellido string
}

func (p *persona) hablar() string {
	return fmt.Sprintf("Hola me llamo %s %s", p.Nombre, p.Apellido)
}

func (p *persona) setNombre(n string) {
	p.Nombre = n
}

type humano interface {
	hablar() string
}

func diAlgo(h humano) {
	fmt.Println(h.hablar())
}

func main() {
	p := persona{
		Nombre:   "David",
		Apellido: "Almaraz",
	}
	// fmt.Println(p.hablar())
	diAlgo(&p)
}
