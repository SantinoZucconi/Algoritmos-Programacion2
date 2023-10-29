package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestApilarAlgunosElementos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("hola")
	pila.Apilar("mundo")
	require.Equal(t, "mundo", pila.Desapilar())
	pila.Apilar("corrector")
	pila.Apilar("como")
	pila.Apilar("va?")
	require.Equal(t, "va?", pila.Desapilar())
	pila.Apilar("estas?")
	require.Equal(t, "estas?", pila.VerTope())
}

func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 10001; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
	}
	require.Equal(t, 10000, pila.VerTope())
	for i := 10000; i >= 0; i-- {
		require.Equal(t, i, pila.Desapilar())
	}
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

func TestDesapilar_Reapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 10; i++ {
		pila.Apilar(i)
	}
	for i := 0; i < 10; i++ {
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	pila.Apilar(10)
	require.Equal(t, 10, pila.VerTope())
	require.False(t, pila.EstaVacia())
}
