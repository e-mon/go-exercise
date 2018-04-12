package main

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	width, height = 600, 320
	iter          = 1 << 14
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill; white; stroke-width: 0.7'"+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < iter; i++ {
		x, y, z := f()
		ax, ay := corner(x+0.1, y, z)
		bx, by := corner(x, y, z)
		cx, cy := corner(x, y+0.1, z)
		dx, dy := corner(x+0.1, y+0.1, z)
		color := colormap(x, y, z)
		fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%06[9]x' stroke='#%06[9]x'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)
	}
	fmt.Println("</svg>")
}

func colormap(x, y, z float64) int {
	min, max := 0x0000ff, 0xff0000
	return int((z/0.5)*float64(max-min) + float64(min))
}

func corner(x, y, z float64) (float64, float64) {
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func ranif(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func f() (float64, float64, float64) {
	t := ranif(1.5*math.Pi, 4.5*math.Pi)

	x := t * math.Cos(t)
	y := t * math.Sin(t)
	z := ranif(0, 0.5)
	return x, y, z
}
