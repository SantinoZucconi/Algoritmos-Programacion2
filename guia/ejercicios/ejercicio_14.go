package ejercicios

import (
	sort "guia/sorts"
	pila "tdas/pila"
)

type Alumno struct {
	Nombre      string
	Parcialito1 int
	Parcialito2 int
	Parcialito3 int
}

func Ejercicio14(l []Alumno) []Alumno {
	l = sort.Counting(l, func(a Alumno) int { return (a.Parcialito1 + a.Parcialito2 + a.Parcialito3) / 3 }, 10)
	l = sort.Counting(l, func(a Alumno) int { return a.Parcialito1 }, 10)
	l = sort.Counting(l, func(a Alumno) int { return a.Parcialito2 }, 10)
	l = sort.Counting(l, func(a Alumno) int { return a.Parcialito3 }, 10)
	pila := pila.CrearPilaDinamica[Alumno]()
	for i := 0; i < len(l); i++ {
		pila.Apilar(l[i])
	}
	for i := 0; i < len(l); i++ {
		l[i] = pila.Desapilar()
	}
	return l
}
