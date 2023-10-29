package ejercicios

// Dados dos arreglos A y B ordenados y sin repeticiones, donde B es igual a A pero
// con un elemento menos, implementar un algoritmo de DyC que permita encontrar el elemento
// faltante de A en B en O(log(n))

func Encontrar_faltante(A []int, B []int) int {
	if len(A) == 1 {
		return A[0]
	}
	medio := (len(A) - 1) / 2
	if A[medio] == B[medio] {
		return Encontrar_faltante(A[medio+1:], B[medio+1:])
	}
	return Encontrar_faltante(A[:medio+1], B[:medio+1])
}
