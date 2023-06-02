package main

import "fmt"

// Devuelve el minimo de dos enteros
func getMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// Representa la tabla de multiplicar del numero indicado entre 1 y 10
func tablaMultiplicar(n int) {
	for i := 0; i <= 9; i++ {
		fmt.Printf("%2d x %2d = %4d\n", i, n, i*n)
	}
}

// Suma los digitos de un numero
func sumaDigits(n int) {
	switch {
	case n < 10:
		return n

	default:

	}
}

func main() {
	fmt.Printf("min (3,4) es %d\n", getMin(3, 4))
	tablaMultiplicar(4)
}
