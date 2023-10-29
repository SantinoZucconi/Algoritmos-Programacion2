package ejercicios

import (
	Pila "tdas/pila"
)

func Ejercicio19(texto string) bool {
	return balanceado(texto)
}

func balanceado(texto string) bool {
	var c rune
	p := Pila.CrearPilaDinamica[string]()
	arr := []string{}
	for _, c = range texto {
		arr = append(arr, string(c))
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == "(" || arr[i] == "[" || arr[i] == "{" {
			p.Apilar(arr[i])
		} else if !p.EstaVacia() {
			if p.VerTope() == contraparte(arr[i]) {
				p.Desapilar()
			} else {
				return false
			}
		}
	}
	if p.EstaVacia() {
		return true
	} else {
		return false
	}

}

func contraparte(t string) string {
	switch t {
	case "(":
		return ")"
	case ")":
		return "("
	case "{":
		return "}"
	case "}":
		return "{"
	case "[":
		return "]"
	case "]":
		return "["
	default:
		panic("no se encontro contraparte")
	}
}
