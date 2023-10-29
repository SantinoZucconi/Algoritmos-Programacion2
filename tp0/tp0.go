package tp0

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	var aux int = *x
	*x = *y
	*y = aux
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	if len(vector) == 0 {
		return -1
	}
	var posicion int = 0
	for i := 0; i < len(vector); i++ {
		if vector[i] > vector[posicion] {
			posicion = i
		}
	}
	return posicion
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func CompararElementos(vector1 []int, vector2 []int, longitud int, valor_mayor int) int {
	for i := 0; i < longitud; i++ {
		if vector1[i] > vector2[i] {
			return 1
		} else if vector1[i] < vector2[i] {
			return -1
		}
	}
	return valor_mayor
}
func Comparar(vector1 []int, vector2 []int) int {
	if len(vector1) > len(vector2) {
		return CompararElementos(vector1, vector2, len(vector2), 1)
	} else if len(vector1) < len(vector2) {
		return CompararElementos(vector1, vector2, len(vector1), -1)
	} else {
		return CompararElementos(vector1, vector2, len(vector1), 0)
	}
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	for i := len(vector) - 1; i > 0; i-- {
		posicion := Maximo(vector[:i+1])
		Swap(&vector[i], &vector[posicion])
	}
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func SumaRecursiva(n int, vector []int) int {
	if n == 1 {
		return vector[0]
	}
	return SumaRecursiva(n-1, vector) + vector[n-1]
}

func Suma(vector []int) int {
	if len(vector) == 0 {
		return 0
	} else {
		return SumaRecursiva(len(vector), vector)
	}
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).

func CadenaCapicuaRecursiva(longitud int, cadena string) bool {
	if longitud == 1 || longitud == 0 {
		return true
	}
	if cadena[longitud-1] == cadena[len(cadena)-longitud] {
		return CadenaCapicuaRecursiva(longitud-1, cadena)
	} else {
		return false
	}
}

func EsCadenaCapicua(cadena string) bool {
	return CadenaCapicuaRecursiva(len(cadena), cadena)
}
