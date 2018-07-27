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

			qx := float64(xmax-xmin) / float64(4*width)
			qy := float64(ymax-ymin) / float64(4*height)
			samples := []color.Color{
				mandlebrot(complex(x-qx, y-qy)),
				mandlebrot(complex(x+qx, y-qy)),
				mandlebrot(complex(x-qx, y+qy)),
				mandlebrot(complex(x+qx, y+qy)),
			}
			img.Set(px, py, mandlebrot(complex(x, y)))
		}
	}

	png.Encode(os.Stdout, img)
}

func mandlebrot(z complex128) color.Color {
	const iter = 120
	var v complex128
	for i := uint8(0); i < iter; i++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - 15*i}
		}
	}
	return color.Black
}

func avgColour(col []color.Color) color.Color {
	var sum [4]float64
	for _, c := range col {
		r, g, b, a := c.RGBA()
		sum[0] += float64(r)
		sum[1] += float64(g)
		sum[2] += float64(b)
		sum[3] += float64(a)
	}
	for i := range sum {
		sum[i] /= float64(len(col))
	}
	return color.RGBA {
		uint8(sum[0]),
		uint8(sum[1]),
		uint8(sum[2]),
		uint8(sum[3]),
	}
}


