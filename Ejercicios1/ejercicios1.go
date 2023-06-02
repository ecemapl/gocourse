package main

import (
	"fmt"
	"math"
)

// Función que devuelve el signo de su argumento
// (-1 si es negativo, +1 si es positivo, y 0 en cualquier otro caso)
func devuelveSigno(param int) string {
	if param > 0 {
		return "+1"
	} else if param < 0 {
		return "-1"
	} else {
		return "0"
	}
}

// Función que, considerando la serie aritmética entre dos números
// Calcula la suma de todos los múltiplos de 3 y de 5 pero que no sean
// múltiplos de 15
func sumaMultiplos3y5no15(a, b int) int {
	suma := 0
	for i := a; i <= b; i++ {
		if i%15 == 0 {
			continue
		} else if i%3 == 0 || i%5 == 0 {
			suma = suma + i
		}
	}
	return suma
}

// Funcion que imprime la tabla de multiplicar del numero que se da como argumento
func imprimeTablaMultiplicar(a int) {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%2d x %2d = %d\n", a, i, a*i)
	}
}

// Funcion que calcula la suma de los digitos de un numero
func sumDigits(number int) int {
	sum := 0

	for number > 0 {
		digit := number % 10
		sum += digit
		number /= 10
	}

	return sum
}

// Funcion que devuelve true de un numero es multiplo de 3
func multiploTres(number int) bool {
	sum := sumDigits(number)

	if sum%3 == 0 {
		return true
	} else {
		return false
	}
}

// Funcion que resuelve ecuacion de segundo grado
func ecuacionSegundoGrado(a, b, c int) (float64, float64) {
	var parteRaiz float64
	var result1, result2 float64

	parteRaiz = math.Sqrt(float64(b*b - 4*a*c))

	result1 = (-float64(b) + parteRaiz) / float64(2*a)
	result2 = (-float64(b) - parteRaiz) / float64(2*a)

	return result1, result2
}

// Funcion que invierte una cadena, por ejemplo "perro" queda transformado en "orrep"
func invierteCadena(input string) string {
	var final = make([]byte, len(input))
	for i, j := 0, len(input)-1; i <= len(input)/2; i, j = i+1, j-1 {
		final[i], final[j] = input[j], input[i]
	}
	return string(final)
}

func main() {
	fmt.Println(devuelveSigno(10))
	fmt.Println(devuelveSigno(-1))
	fmt.Println(devuelveSigno(0))

	// Imprime suma de multiplos de 3 y de 5 pero no de 15 entre 0 y 300000
	fmt.Printf("La suma de multiplos de 3 y 5, pero no de 15, entre %d y %d es %d\n", 0, 300000, sumaMultiplos3y5no15(0, 300000))

	// Imprime la tabla de multiplicar del argumento
	imprimeTablaMultiplicar(7)

	// Determina si un numero es multiplo de 3, para ello la suma de sus cifras tiene que ser multiplo de 3
	// 128 no es multiplo de 3
	if multiploTres(128) {
		fmt.Println("128 es multiplo de 3")
	} else {
		fmt.Println("128 no es multiplo de 3")
	}
	// 129 si es multiplo de 3
	if multiploTres(129) {
		fmt.Println("129 es multiplo de 3")
	} else {
		fmt.Println("129 no es multiplo de 3")
	}

	// Resuelve ecuacion segundo grado
	a, b, c := 2, 4, 1
	d, e := ecuacionSegundoGrado(a, b, c)
	fmt.Printf("Las soluciones de %dx^2+%dx+%d son %f y %f\n", a, b, c, d, e)

	// Determina si una cadena es un palindromo
	var cadena string = "Hola!"
	fmt.Println(invierteCadena(cadena))
	if cadena == invierteCadena(cadena) {
		fmt.Printf("%s es un palindromo\n", cadena)
	} else {
		fmt.Printf("%s no es un palindromo\n", cadena)
	}
}
