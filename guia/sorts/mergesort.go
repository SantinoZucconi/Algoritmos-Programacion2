package guia

func Merge(lista []int) []int {
	if len(lista) == 1 {
		return lista
	}
	medio := len(lista) / 2
	arreglo_izq := lista[:medio]
	arreglo_der := lista[medio:]
	izq_ordenada := Merge(arreglo_izq)
	der_ordenada := Merge(arreglo_der)
	arreglo_ordenado := intercalar_arreglo(izq_ordenada, der_ordenada, len(izq_ordenada)+len(der_ordenada))
	return arreglo_ordenado
}

func intercalar_arreglo(arr1 []int, arr2 []int, long int) []int {
	arreglo_intercalado := make([]int, len(arr1)+len(arr2))
	for i := 0; i < long; i++ {
		if len(arr2) == 0 {
			arr1 = intercalar(&arreglo_intercalado[i], arr1)
		} else if len(arr1) == 0 {
			arr2 = intercalar(&arreglo_intercalado[i], arr2)
		} else {
			if arr1[0] < arr2[0] {
				arr1 = intercalar(&arreglo_intercalado[i], arr1)
			} else if arr1[0] >= arr2[0] {
				arr2 = intercalar(&arreglo_intercalado[i], arr2)
			}
		}
	}
	return arreglo_intercalado
}

func intercalar(arr *int, arr_intercalar []int) []int {
	*arr = arr_intercalar[0]
	arr_intercalar = arr_intercalar[1:]
	return arr_intercalar
}
