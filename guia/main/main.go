package main

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
)

func main() {
	abb := TDADiccionario.CrearAbbPrueba[int, int](func(a, b int) int { return (b - a) })
	abb.Guardar(8, 10)
	abb.Guardar(4, 2)
	abb.Guardar(10, 12)
	abb.Guardar(9, 12)
	abb.Guardar(11, 12)
	abb.Guardar(3, 12)
	abb.Guardar(5, 12)

	fmt.Println(abb.Altura())
}
