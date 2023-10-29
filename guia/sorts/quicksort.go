package guia

func Quick(arr []int) []int {
	if len(arr) == 1 || len(arr) == 0 {
		return arr
	}
	pivote := arr[0]
	arr_menores := make([]int, len(arr)-1)
	arr_mayores := make([]int, len(arr)-1)
	ptr_menores := 0
	ptr_mayores := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > pivote {
			arr_mayores[ptr_mayores] = arr[i]
			ptr_mayores += 1
		} else {
			arr_menores[ptr_menores] = arr[i]
			ptr_menores += 1
		}
	}
	arr_mayores = arr_mayores[:ptr_mayores]
	arr_menores = arr_menores[:ptr_menores]
	arr_mayores = Quick(arr_mayores)
	arr_menores = Quick(arr_menores)
	arreglo_ordenado := append(arr_menores, pivote)
	arreglo_ordenado = append(arreglo_ordenado, arr_mayores...)
	return arreglo_ordenado
}
