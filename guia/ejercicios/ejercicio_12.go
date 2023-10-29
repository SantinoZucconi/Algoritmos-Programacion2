package ejercicios

func Ejercicio12(n int) int {
	return aux(n, 0, n)
}

func aux(n, min, max int) int {
	if min == max {
		return min
	}
	medio := (min + max) / 2
	if medio*medio <= n && (medio+1)*(medio+1) > n {
		return medio
	}
	if medio*medio > n {
		return aux(n, min, medio)
	}
	return aux(n, medio+1, max)
}
