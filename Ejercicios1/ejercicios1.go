package main

import "fmt"

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
}
