package ejercicios

func Ejercicio11(n []int) int {
	return auxs(n, 0, len(n))
}

func auxs(n []int, min int, max int) int {
	if min == max {
		return min
	}
	medio := (min + max) / 2
	if n[medio] == 1 {
		return auxs(n, medio+1, max)
	} else {
		return auxs(n, min, medio)
	}
}
