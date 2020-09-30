// lissajous uses standard image packages from the go library to create a sequence of bit-mapped images and then encode
// the sequence as a gif animation;

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

// composite literal;
// a compact notation for instantiating any of go's composite types from a sequence of element values;
var palette = []color.Color{color.White, color.Black} // slice

const (
	whiteIndex = 0 // first color in the palette
	blackIndex = 1 // next color in the palette
)

func main() {

	lissajous(os.Stdout)

}

func lissajous(output io.Writer) {

	const (
		cycles     = 5     // number of complete x oscillator revolutions
		resolution = 0.001 // angular resolution
		size       = 100   // image canvas covers
		nframes    = 64    // number of animation frames
		delay      = 8     // delay between frames in 10ms units
	)

	frequency := rand.Float64() * 3.0 // relative frequency of the y oscillator

	// the variable anim is a struct of type gif.GIF;
	anim := gif.GIF{LoopCount: nframes} // composite literal

	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {

		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += resolution {
			x := math.Sin(t)
			y := math.Sin(t*frequency + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}

		phase += 0.1

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

	}

	gif.EncodeAll(output, &anim) // NOTE: we are ignoring encoding errors

}
