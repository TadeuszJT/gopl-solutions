// Modify the Lissajous program to produce images in multiple colours by
// adding more values to palette and then displaying them by changing the
// third argument of Set-ColorIndex in some interesting way.
//
// Solved by Tadeusz Tomoszek 12.7.2018

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.RGBA{0xff, 0xff, 0xff, 0xff}, // White
	color.RGBA{0x00, 0x00, 0x00, 0xff}, // Black
	color.RGBA{0xff, 0x00, 0x00, 0xff}, // Red
	color.RGBA{0xff, 0xff, 0x00, 0xff}, // Yellow
	color.RGBA{0x00, 0xff, 0x00, 0xff}, // Green
	color.RGBA{0x00, 0xff, 0xff, 0xff}, // Cyan
	color.RGBA{0x00, 0x00, 0xff, 0xff}, // Blue
	color.RGBA{0xff, 0x00, 0xff, 0xff}, // Purple


}

const (
	whiteIndex = 0 // first colour in palette
	blackIndex = 1 // next colour in palette
	redIndex = 2
	purpleIndex = 7
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		var index uint8 = redIndex
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				index)

			index++
			if index > purpleIndex {
				index = redIndex
			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Ignoring encoding errors
}

