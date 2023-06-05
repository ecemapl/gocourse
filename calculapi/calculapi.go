package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var n uint64 = 500000000 //Número de puntos
	var nPuntosDentro uint64
	var pi float64

	pts := make(plotter.XYs, 1000)
	// Create a new plot, set its title and
	// axis labels.
	p := plot.New()

	//p.Title.Text = "Points Example"
	//p.X.Label.Text = "X"
	//p.Y.Label.Text = "Y"
	// Draw a grid behind the data
	//p.Add(plotter.NewGrid())

	var i uint64
	for ; i < n; i++ {
		x := r.Float64()
		y := r.Float64()

		if x*x+y*y <= 1.0 {
			nPuntosDentro++
		}
		if i < 1000 {
			pts[i].X = x
			pts[i].Y = y
		}

	}
	pi = float64(nPuntosDentro) * 4.0 / float64(n)

	fmt.Printf("Pi = %f\n", pi)

	// Make a scatter plotter and set its style.
	s, err := plotter.NewScatter(pts)
	if err != nil {
		panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	p.Add(s)

	// Save the plot to a PNG file.
	if err := p.Save(15*vg.Centimeter, 15*vg.Centimeter, "scatter.png"); err != nil {
		panic(err)
	}

}
