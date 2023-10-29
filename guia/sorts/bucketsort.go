package guia

func Bucket(arr []int, rangos []int) []int {
	lista_ordenada := make([]int, 0)
	listas_baldes := make([][]int, len(rangos))
	for j := 0; j < len(arr); j++ {
		for i := 0; i < len(rangos); i++ {
			if arr[j] <= rangos[i] {
				listas_baldes[i] = append(listas_baldes[i], arr[j])
				i += len(rangos)
			}
		}
	}
	for i := 0; i < len(listas_baldes); i++ {
		Quick(listas_baldes[i])
		lista_ordenada = append(lista_ordenada, listas_baldes[i]...)
	}
	return lista_ordenada
}
