package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestEncolarAlgunosElementos(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("hola")
	require.Equal(t, "hola", cola.Desencolar())
	cola.Encolar("hola")
	cola.Encolar("como")
	cola.Encolar("va")
	require.Equal(t, "hola", cola.Desencolar())
	require.Equal(t, "como", cola.Desencolar())
	require.Equal(t, "va", cola.Desencolar())
}

func TestVolumenCola(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 10001; i++ {
		cola.Encolar(i)
	}
	for i := 0; i < 10000; i++ {
		require.Equal(t, i, cola.Desencolar())
	}
	require.Equal(t, 10000, cola.VerPrimero())
	cola.Desencolar()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestEncolaryDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 100; i++ {
		cola.Encolar(i)
	}
	for i := 0; i < 100; i++ {
		require.Equal(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
	cola.Encolar(10)
	require.Equal(t, 10, cola.VerPrimero())
	cola.Encolar(11)
	require.Equal(t, 10, cola.VerPrimero())
	cola.Desencolar()
	require.Equal(t, 11, cola.VerPrimero())
	cola.Desencolar()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}
