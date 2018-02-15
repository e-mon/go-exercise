package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

const (
	host = "localhost"
	port = "8000"
)

func main() {
	fmt.Printf("listen %s:%s..\n", host, port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	m, _ := url.ParseQuery(r.URL.RawQuery)
	cycles := 5

	if val, ok := m["cycles"]; ok {
		n, err := strconv.Atoi(val[0])
		if err == nil {
			cycles = n
		}
	}

	lisaajous(w, float64(cycles))
}

func lisaajous(out io.Writer, cycles float64) {
	const (
		res       = 0.001
		size      = 100
		nframes   = 64
		delay     = 8
		max_color = 4
	)

	var palette = []color.Color{color.Black}

	for i := 0; i < max_color; i++ {
		palette = append(palette, color.RGBA{uint8(rand.Intn(0xff)),
			uint8(rand.Intn(0xff)), uint8(rand.Intn(0xff)), 0xff})
	}

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(max_color-1)+1))
		}
		phase += 1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
