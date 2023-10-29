package ejercicios

func Encontrar_joya(arr []int) int {
	return encontrar_joya_aux(arr, 0, len(arr))
}

func encontrar_joya_aux(arr []int, min, max int) int {
	if min == max {
		return min
	}
	medio := (min + max) / 2
	if balanza(arr[:medio], arr[medio:]) == 1 {
		return encontrar_joya_aux(arr, min, medio)
	} else {
		return encontrar_joya_aux(arr, medio+1, max)
	}
}

func balanza(a1 []int, a2 []int) int {
	suma1 := maximo(a1)
	suma2 := maximo(a2)
	if suma1 > suma2 {
		return 1
	} else if suma1 < suma2 {
		return -1
	} else {
		return 0
	}
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
