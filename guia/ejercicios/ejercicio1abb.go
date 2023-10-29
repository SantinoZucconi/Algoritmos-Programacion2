package ejercicios

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
)

func Ejercicio1ABB() {
	abb := TDADiccionario.CrearAbbPrueba[int, int](func(a, b int) int { return (b - a) })
	abb.Guardar(5, 6)
	abb.Guardar(10, 6)
	abb.Guardar(3, 6)
	abb.Guardar(1, 6)
	abb.Guardar(15, 6)
	abb.Guardar(13, 6)
	abb.Guardar(11, 6)
	fmt.Println(abb.Altura())
}
