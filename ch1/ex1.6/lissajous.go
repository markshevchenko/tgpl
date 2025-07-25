package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0x0F, 0x0F, 0x0F, 0xFF},
	color.RGBA{0x1F, 0x1F, 0x1F, 0xFF},
	color.RGBA{0x2F, 0x2F, 0x2F, 0xFF},
	color.RGBA{0x3F, 0x3F, 0x3F, 0xFF},
	color.RGBA{0x4F, 0x4F, 0x4F, 0xFF},
	color.RGBA{0x5F, 0x5F, 0x5F, 0xFF},
	color.RGBA{0x6F, 0x6F, 0x6F, 0xFF},
	color.RGBA{0x7F, 0x7F, 0x7F, 0xFF},
	color.RGBA{0x8F, 0x8F, 0x8F, 0xFF},
	color.RGBA{0x9F, 0x9F, 0x9F, 0xFF},
	color.RGBA{0xAF, 0xAF, 0xAF, 0xFF},
	color.RGBA{0xBF, 0xBF, 0xBF, 0xFF},
	color.RGBA{0xCF, 0xCF, 0xCF, 0xFF},
	color.RGBA{0xDF, 0xDF, 0xDF, 0xFF},
	color.RGBA{0xEF, 0xEF, 0xEF, 0xFF},
	color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
}

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

	rand.Seed(time.Now().UTC().UnixNano())

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		var colorIndex uint8 = 1

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size*0.5), size+int(y*size+0.5), colorIndex)

			colorIndex++
			if int(colorIndex) == len(palette) {
				colorIndex = 1
			}
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
