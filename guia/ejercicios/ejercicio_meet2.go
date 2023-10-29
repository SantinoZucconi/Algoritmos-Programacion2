package ejercicios

// Implementar un algoritmo de ordenamiento, que sea lineal,
// que permita definir el orden en una fila de personas para comprar
// una Cajita CampeónFeliz en un establecimiento de comida rápida.
// Los datos (structs) a ordenar cuentan con edad (número), nombre (string) y
// nacionalidad (enumerativo, de 32 valores posibles).
// Primero deben ir los niños (todos con edad menor o igual a 12),
// y estos deben ordenarse por edad (de menor a mayor),
// independientemente de la nacionalidad. Luego deben ir los “no niños”,
// que primero deben estar ordenados por nacionalidad (segundo Francia, por ejemplo)
// y en caso de igualdad de nacionalidad, por edad, también de menor a mayor.
// En caso de necesitar algún ordenamiento auxiliar
// (en caso que una parte del algoritmo implique hacer BucketSort o RadixSort),
// puede considerarse como ya implementado, pero justificar (en caso de necesitar justificación)
// por qué puede utilizarse dicho algoritmo dadas las condiciones del problema.
// Justificar la complejidad de la función implementada.
// El desarrollo de la complejidad debe estar completo para el problema en cuestión,
// no se aceptarán resultados parciales genéricos.

import sort "guia/sorts"

type Persona struct {
	edad         int
	nombre       string
	nacionalidad int
}

func CajitaCameponFeliz(personas []Persona) []Persona {
	personas_Mayores := make([]Persona, 0)
	personas_Menores := make([]Persona, 0)
	for _, persona := range personas {
		if persona.edad <= 12 {
			personas_Menores = append(personas_Menores, persona)
		} else {
			personas_Mayores = append(personas_Mayores, persona)
		}
	}
	personas_Menores = sort.Counting[Persona](personas_Menores, func(p Persona) int { return p.edad }, 12)
	personas_Mayores = sort.Counting[Persona](personas_Mayores, func(p Persona) int { return p.edad }, 100)
	personas_Mayores = sort.Counting[Persona](personas_Mayores, func(p Persona) int { return p.nacionalidad }, 32)
	total_personas := make([]Persona, 0)
	total_personas = append(total_personas, personas_Menores...)
	total_personas = append(total_personas, personas_Mayores...)
	return total_personas
}
