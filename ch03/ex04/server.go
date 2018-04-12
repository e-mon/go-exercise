package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

const (
	host    = "localhost"
	port    = "8000"
	xyrange = 30.0
	angle   = math.Pi / 6
)

type params struct {
	Height int
	Width  int
	Iter   int
}

func (p *params) set(values url.Values) {
	setDefault := func(key string, defaultValue int) int {
		v, err := strconv.Atoi(values.Get(key))
		if err != nil {
			return defaultValue
		}
		return v
	}

	p.Height = setDefault("height", 600)
	p.Width = setDefault("width", 320)
	p.Iter = setDefault("iter", 1<<10)

}

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("listen %s:%s..\n", host, port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	p := &params{}
	p.set(r.URL.Query())
	swissrole(w, p)
}

func swissrole(out io.Writer, p *params) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill; white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", p.Width, p.Height)
	for i := 0; i < p.Iter; i++ {
		x, y, z := f()
		ax, ay := corner(x+0.1, y, z, p)
		bx, by := corner(x, y, z, p)
		cx, cy := corner(x, y+0.1, z, p)
		dx, dy := corner(x+0.1, y+0.1, z, p)
		color := colormap(x, y, z)
		fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%06[9]x' stroke='#%06[9]x' />\n", ax, ay, bx, by, cx, cy, dx, dy, color)
	}
	fmt.Fprintln(out, "</svg>")
}

func colormap(x, y, z float64) int {
	min, max := 0x0000ff, 0xff0000
	return int((z/0.5)*float64(max-min) + float64(min))
}

func corner(x, y, z float64, p *params) (float64, float64) {
	xyscale := float64(p.Width) / 2 / xyrange
	zscale := float64(p.Height) * 0.4
	sx := float64(p.Width)/2 + (x-y)*cos30*xyscale
	sy := float64(p.Height)/2 + (x+y)*sin30*xyscale - z*zscale
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
