package cola

type colaEnlazada[T any] struct {
	inicio *nodo[T]
	fin    *nodo[T]
}

type nodo[T any] struct {
	valor     T
	siguiente *nodo[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{}
}

func (c *colaEnlazada[T]) Encolar(dato T) {
	nuevoNodo := &nodo[T]{dato, nil}
	if c.EstaVacia() {
		c.inicio = nuevoNodo
	} else {
		c.fin.siguiente = nuevoNodo
	}
	c.fin = nuevoNodo
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	valor := c.inicio.valor
	c.inicio = c.inicio.siguiente
	if c.inicio == nil {
		c.fin = nil
	}
	return valor

}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.inicio == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.inicio.valor
}
