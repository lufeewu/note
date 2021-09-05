package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"reflect"

	"github.com/sirupsen/logrus"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func key_words() {
	var condition string
	app := kingpin.New("name", "a key words")
	app.Flag("condition", "switch condition").Default("test").StringVar(&condition)
	app.Parse(os.Args[1:])
	// logrus.Infoln(os.Args[1:])

	switch condition {
	case "test":
		logrus.Infoln("test")
		fallthrough
	case "test2":
		logrus.Infoln("test2")
		fallthrough
	default:
		logrus.Infoln("default")
	case "test3":
		logrus.Infoln("test3")
		break // break "switch" "select" "for"(loop)
	}

	// rune unitptr
	var r rune
	type cc uintptr
	var c cc = 0xffffffffffffffff
	logrus.Infoln(c)
	logrus.Infoln(r)
	str := "中国话"
	for k, runeValue := range str {
		logrus.Infof("%#U starts at byte position %d\n", runeValue, k)
		logrus.Infoln(reflect.TypeOf(runeValue))
	}

	var com complex64 = 1 + 21i
	logrus.Infoln(real(com), imag(com))

	var m map[int]int
	m = make(map[int]int)
	m[3] = 4
	m[4] = 5
	logrus.Infoln(m, len(m))
	logrus.Infoln(m[4])
	delete(m, 3)
	logrus.Infoln(m, m[3], m[4], m[5], len(m))
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	file, _ := os.Create("./tmp.png")
	png.Encode(file, img) // NOTE: ignoring errors
	defer file.Close()
}
