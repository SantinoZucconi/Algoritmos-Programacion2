package guia

func Counting[T any](lista []T, criterio func(T) int, max_rango int) []T {
	contadores := make([]int, max_rango+1)
	for _, elemento := range lista {
		contadores[criterio(elemento)] += 1
	}
	sumas_acumuladas := make([]int, max_rango+1)
	for i := 0; i < len(contadores)-1; i++ {
		sumas_acumuladas[i+1] += sumas_acumuladas[i] + contadores[i]
	}
	nuevo_arreglo := make([]T, len(lista))
	for i := 0; i < len(lista); i++ {
		nuevo_arreglo[sumas_acumuladas[criterio(lista[i])]] = lista[i]
		sumas_acumuladas[criterio(lista[i])] += 1
	}
	return nuevo_arreglo
}
