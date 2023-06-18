package bitmap

import (
	"fmt"
	"math"
	"strings"
)

// Función que devuelve un string lo suficientemente largo para alojar length bits
// Sugerencia: usa la función "strings.Repeat" (mira la documentación) para replicar un string tantas veces como quieras
func NewBitmap(length int) string {
	numBytes := 0

	if length%8 > 0 {
		numBytes = int(length/8) + 1
	} else {
		numBytes = int(length / 8)
	}

	//funcionará esto?
	//var bitmap string = make([]byte, numBytes)

	//Ha sido raro para inicializar a 0 (sacado de https://golangdocs.com/strings-in-golang)
	//return strings.Repeat(string([]rune{0x0000}), numBytes) // Esta inicializa a cero de verdad
	return strings.Repeat("0", numBytes) // Estas es para probar porque "0" se mapea a un 48 en binario

}

// devuelve el valor del bit en la posición i del bitmap almacenado en el string entregado en primer lugar
func ReadBit(bitmap string, i int) byte {
	var resultado byte = 0

	//Primero obtenemos el byte donde está el i-esimo bit
	var bytePos int = int(i / 8)

	//Ahora el bit dentro del byte
	var bitPos int = i % 8

	//Ahora la máscara para extraer el bit
	var mask byte = byte(math.Pow(float64(2), float64(7-bitPos)))

	var t4 = make([]byte, len(bitmap))
	copy(t4, bitmap)
	resultado = t4[bytePos] & mask

	if resultado > 0 {
		return '1'
	} else {
		return '0'
	}
}

// devuelve un bitmap que es una copia del entregado como primer parámetro y que, además, tiene activado el bit i-ésimo
//func EnableBit(bitmap stirng, i int) string {

//}

// devuelve un bitmap que es una copia del entregado como primer parámetro y que, además, tiene desactivado el bit i-ésimo
//func DisableBit(bitmap string, i int) string {

//}

// imprime por pantalla los contenidos de todos los bits del bitmap entregado como argumento
func PrintBitmap(bitmap string) {

	var t4 = make([]byte, len(bitmap))
	copy(t4, bitmap)

	for i := 0; i <= len(bitmap)-1; i++ {
		for j := 0; j <= 7; j++ {
			if t4[i]&byte(math.Pow(float64(2), float64(7-j))) > 0 {
				fmt.Println("1")
			} else {
				fmt.Println("0")
			}
		}
	}
}
