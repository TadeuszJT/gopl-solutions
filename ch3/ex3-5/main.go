package main

import (
	"os"
	"math/cmplx"
	"image"
	"image/png"
	"image/color"
)

var palette = []color.Color {
	color.RGBA{148, 0, 211, 255},
	color.RGBA{75, 0, 130, 255},
	color.RGBA{0, 0, 255, 255},
	color.RGBA{0, 255, 0, 255},
	color.RGBA{255, 255, 0, 255},
	color.RGBA{255, 127, 0, 255},
	color.RGBA{255, 0, 0, 255},
}

func main() {
	const (
		width = 800
		height = 800
		xmin = -2
		xmax = 2
		ymin = -2
		ymax = 2
	)
	rec := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rec)

	for px := 0; px < width; px++ {
		x := float64(px)/width * (xmax-xmin) + xmin
		for py := 0; py < height; py++ {
			y := float64(py)/height * (ymax-ymin) + ymin
			img.Set(px, py, mandlebrot(complex(x, y)))
		}
	}

	png.Encode(os.Stdout, img)
}

func mandlebrot(z complex128) color.Color {
	const iter = 200
	var v complex128
	for i := 0; i < iter; i++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette[i % len(palette)]
		}
	}
	return color.Black
}

