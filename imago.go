// https://www.socketloop.com/tutorials/golang-get-rgba-values-of-each-image-pixel
// https://stackoverflow.com/questions/33186783/get-a-pixel-array-from-from-golang-image-image
// https://socketloop.com/tutorials/golang-convert-integer-to-binary-octal-hexadecimal-and-back-to-integer
// https://dev.to/andyhaskell/how-i-made-a-slick-personal-logo-with-go-s-standard-library-29j9
package main

import (
	//"fmt"
	"log"
	"flag"
	"image"
	"image/color"
	"image/png"
	// ^ https://golang.org/pkg/image/
	"image/jpeg"
	"os"
)

func decodeImg(myImage string) (image.Image, string, error) {
	file, err := os.Open(myImage)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	data, format, err := image.Decode(file)
	if err != nil {
		return nil, "", err
	}
	return data, format, nil
}

func brightness(r, g, b, a uint8, number int) (uint8, uint8, uint8, uint8) {
	// Converting a number to uint8 it does number % 255
	// This is really ugly but i have to check if the value 
	// is over 255 so i can manually cap it so it doesnt
	// loop around.

	// TODO: maybe (number - (number - 255)) % 255
	if int(r) + number > 255 {
		r = 255
	} else if int(r) + number < 0 {
		r = 0
	} else {
		r += uint8(number)
	}

	if int(g) + number > 255 {
		g = 255
	} else if int(g) + number < 0 {
		g = 0
	} else {
		g += uint8(number)
	}

	if int(b) + number > 255 {
		b = 255
	} else if int(b) + number < 0 {
		b = 0
	} else {
		b += uint8(number)
	}

	return r, g, b, a
}

func randomizer(r, g, b, a uint8, number int) (uint8, uint8, uint8, uint8) {
	r *= uint8(number)
	g *= uint8(number)
	b *= uint8(number)
	return r, g, b, a
}

func inverse(r, g, b, a uint8) (uint8, uint8, uint8, uint8) {
	r = 255 - r
	g = 255 - g
	b = 255 - b
	return r, g, b, a
}

func greyscale(r, g, b, a uint8) (uint8, uint8, uint8, uint8) {
	// https://github.com/smtnsk/go-course/blob/master/assignments/project/greyscaler.go
	grey := uint8((float64(r) + float64(g) + float64(b)) / 3)
	return grey, grey, grey, a
}

func recolor(data image.Image, bounds image.Rectangle, number int, mode string) *image.RGBA {
	recolored := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
		pixel := data.At(x, y)
		pixelRGBA := color.RGBAModel.Convert(pixel).(color.RGBA)
		r := pixelRGBA.R
		g := pixelRGBA.G
		b := pixelRGBA.B
		a := pixelRGBA.A
		if mode == "inverse" {
			r, g, b, a = inverse(r, g, b, a)
		} else if mode == "greyscale" {
			r, g, b, a = greyscale(r, g, b, a)
		} else if mode == "brightness" {
			r, g, b, a = brightness(r, g, b, a, number)
		} else if number != 0 {
			r, g, b, a = randomizer(r, g, b, a, number)
		}
		newColor := color.RGBA{R: r, G: g, B: b, A: a}
		recolored.Set(x, y, newColor)
		}
	}
	return recolored
}

func main() {

	var image string
	var output string
	var number int
	var mode string

	// TODO: make better help text...
	flag.StringVar(&image, "f", "", "File to use")
	flag.StringVar(&output, "o", "", "Output file")
	flag.IntVar(&number, "n", 0, "Value number. if mode is omitted this does magic to images")
	flag.StringVar(&mode, "m", "", "Edit mode (eg. Inverse, grayscale)\nCurrently supported:\ninverse\ngreyscale\nbrightness (needs a value using -n (-255 - 255))")
	flag.Parse()

	// Decode image into something useable and define bounds
	data, format, err := decodeImg(image)
	if err != nil {
		log.Fatal(err)
	}
	bounds := data.Bounds()
	
	// Create an outfile
	out, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	
	// Make the new image
	recolored := recolor(data, bounds, number, mode)
	if format == "png" {
		if err := png.Encode(out, recolored); err != nil {
		log.Fatal(err)
		}
	} else if format == "jpeg" {
		if err := jpeg.Encode(out, recolored, nil); err != nil {
		log.Fatal(err)
		}
	}
}