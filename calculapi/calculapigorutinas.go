package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	//var numCores int = runtime.NumCPU()
	var numCores int = 4
	var n uint64 = 500000000 //Número de puntos por gorutina
	var nPuntosDentro = make([]uint64, numCores)
	var nPuntosDentroTotal uint64
	var pi float64
	var wg sync.WaitGroup
	wg.Add(numCores)

	start := time.Now()

	for j := 0; j < numCores; j++ {
		go func(a *uint64) {
			defer wg.Done()
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			var i uint64
			for ; i < n; i++ {
				x := r.Float64()
				y := r.Float64()

				if x*x+y*y <= 1.0 {
					(*a)++
				}
			}
		}(&nPuntosDentro[j])

	}
	wg.Wait()

	for j := 0; j < numCores; j++ {
		nPuntosDentroTotal += nPuntosDentro[j]
	}

	elapsed := time.Since(start)

	pi = float64(nPuntosDentroTotal/uint64(numCores)) * 4.0 / float64(n)

	fmt.Printf("Numero de cores utilizados = %d\n", numCores)
	fmt.Printf("Tiempo invertido: %s\n", elapsed)
	fmt.Printf("Pi = %f\n", pi)

}
