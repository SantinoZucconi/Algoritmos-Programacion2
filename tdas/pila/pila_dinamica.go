package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{datos: make([]T, 1), cantidad: 0}
}

func (p *pilaDinamica[T]) Apilar(nuevo_dato T) {
	if p.cantidad == len(p.datos) {
		p.redimensionar(len(p.datos) * 2)
	}
	p.datos[p.cantidad] = nuevo_dato
	p.cantidad += 1
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	} else {
		dato_desapilado := p.datos[p.cantidad-1]
		p.cantidad -= 1
		if p.cantidad == (len(p.datos)/4) && !p.EstaVacia() {
			p.redimensionar(len(p.datos) / 2)
		}
		return dato_desapilado
	}
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	} else {
		return p.datos[p.cantidad-1]
	}
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	if p.cantidad == 0 {
		return true
	} else {
		return false
	}
}

func (p *pilaDinamica[T]) redimensionar(tamaño int) {
	nuevo_arreglo := make([]T, tamaño)
	copy(nuevo_arreglo, p.datos)
	p.datos = nuevo_arreglo
}
