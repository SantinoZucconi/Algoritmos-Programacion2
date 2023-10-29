package guia

import TDACola "tdas/cola"

func Radix1(arr []int) []int {
	total_digitos := 1
	digitos := maximo(arr)
	for i := 1; digitos/i > 0; i *= 10 {
		total_digitos *= 10
	}
	for k := 1; total_digitos/k > 0; k *= 10 {
		lista := make([]TDACola.Cola[int], 10)
		for i := 0; i < 10; i++ {
			lista[i] = TDACola.CrearColaEnlazada[int]()
		}
		for _, elementos := range arr {
			lista[elementos%(10*k)/k].Encolar(elementos)
		}
		indice := 0
		for j := 0; j < len(lista); j++ {
			for !lista[j].EstaVacia() {
				arr[indice] = lista[j].Desencolar()
				indice++
			}
		}
	}
	return arr
}

func Radix2(arr []int) []int {
	cant_digitos := 0
	total_digitos := 1
	digitos := maximo(arr)
	for i := 1; digitos/i > 0; i *= 10 {
		cant_digitos++
		total_digitos *= 10
	}
	total_digitos /= 10
	for k := 1; total_digitos/k > 0; k *= 10 {
		arr = Counting[int](arr, func(num int) int { return num % (10 * k) / k }, 10)
	}
	return arr
}

func maximo(vector []int) int {
	if len(vector) == 0 {
		return -1
	}
	var posicion int = 0
	for i := 0; i < len(vector); i++ {
		if vector[i] > vector[posicion] {
			posicion = i
		}
	}
	return vector[posicion]
}
