package ejercicios

func Ejercicio8(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	medio := len(arr) / 2
	arr_izq := arr[0:medio]
	arr_der := arr[medio:]
	min_izquierda := Ejercicio8(arr_izq)
	min_derecha := Ejercicio8(arr_der)
	if min_izquierda < min_derecha {
		return min_izquierda
	} else {
		return min_derecha
	}
}
