package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

func main() {
	// Definir flags
	lFlag := flag.Bool("l", false, "Mostrar detalles extendidos")
	tFlag := flag.Bool("t", false, "Ordenar por fecha de modificación")
	rFlag := flag.Bool("r", false, "Ordenar en orden inverso")
	flag.Parse()

	// Obtener el directorio a mostrar
	dir := "."
	args := flag.Args()
	if len(args) > 0 {
		dir = args[0]
	}

	// Comprobación de permisos de lectura
	_, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("No se pueden leer los archivos en el directorio %s: %s", dir, err)
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// Aplicar los flags
	if *lFlag {
		// Mostrar detalles extendidos
		for _, file := range files {
			info := file.Mode().String()
			size := file.Size()
			modTime := file.ModTime().Format("Jan 02 15:04")
			name := file.Name()
			fmt.Printf("%s\t%d\t%s\t%s\n", info, size, modTime, name)
		}
	} else if *tFlag {
		// Ordenar por fecha de modificación
		sort.Slice(files, func(i, j int) bool {
			return files[i].ModTime().After(files[j].ModTime())
		})
	} else if *rFlag {
		// Ordenar en orden inverso
		for i := len(files)/2 - 1; i >= 0; i-- {
			opp := len(files) - 1 - i
			files[i], files[opp] = files[opp], files[i]
		}
	} else {
		// Mostrar los nombres de archivos
		for _, file := range files {
			fmt.Println(file.Name())
		}
	}
}
