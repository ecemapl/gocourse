package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Por favor, proporciona el nombre del programa o archivo como argumento.")
		return
	}

	name := os.Args[1] // Obtener el nombre del programa o archivo del segundo argumento

	path, err := exec.LookPath(name)
	if err == nil {
		info, err := os.Stat(path)
		if err != nil {
			fmt.Printf("No se pudo obtener información del programa o archivo %s.\n", name)
			return
		}

		mode := info.Mode()
		if mode.IsRegular() && (mode&0111) != 0 {
			fmt.Printf("El programa %s es ejecutable y se encuentra en la ruta: %s\n", name, path)
		} else {
			fmt.Printf("El programa %s no es ejecutable o no se encuentra en la ruta especificada.\n", name)
		}
	} else {
		fmt.Printf("No se encontró el programa o archivo %s en el sistema.\n", name)
	}
}
