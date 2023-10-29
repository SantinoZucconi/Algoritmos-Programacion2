package ejercicios

func Ejercicio9(arr []int) bool {
	condicion := true
	esta_ordenado(arr, &condicion)
	return condicion
}

func esta_ordenado(arr []int, c *bool) []int {
	if len(arr) == 1 {
		return arr
	}
	medio := len(arr) / 2
	izq := esta_ordenado(arr[0:medio], c)
	der := esta_ordenado(arr[medio:], c)
	if izq[len(izq)-1] > der[0] {
		*c = false
	}
	return arr
} // T = 2T(n/2) + O(1) = O(n)
