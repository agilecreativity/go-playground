package main

import (
	"flag"
	"image"
	_ "image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

var (
	jpgOut = flag.String("jpg", "", "output to a jpeg")
	pngOut = flag.String("png", "", "output to a png")
	in     = flag.String("in", "", "input file like *.gif")
)

type encf func(io.Writer, image.Image) error

func encode(encoder encf, img image.Image, filename string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Printf("failed to open :%s with error: %s", filename, err)
		return
	}

	defer file.Close()

	err = encoder(file, img)

	if err != nil {
		log.Printf("failed to encode :%s, with error: %s", filename, err)
	}
}

func jpegEncode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, &jpeg.Options{Quality: 80})
}

func decode(filename string) image.Image {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Failed to open file: %s, error: %s", filename, err)
	}

	defer file.Close()

	img, _, err := image.Decode(file)

	if err != nil {
		log.Fatalf("Error decoding image: %s", err)
	}
	return img
}

func main() {
	flag.Parse()

	img := decode(*in)

	if *pngOut != "" {
		encode(png.Encode, img, *pngOut)
	}

	if *jpgOut != "" {
		encode(jpegEncode, img, *jpgOut)
	}
}
