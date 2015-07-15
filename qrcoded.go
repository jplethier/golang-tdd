package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello QR Code")
	file, err := os.Create("qrcode.png")
	if err != nil {
		LogError(err)
	}
	defer file.Close()

	err = GenerateQRCode(file, "555-2368", Version(1))
	if err != nil {
		LogError(err)
	}
}

func GenerateQRCode(w io.Writer, code string, version Version) error {
	size := version.PatternSize()
	img := image.NewNRGBA(image.Rect(0, 0, size, size))
	return png.Encode(w, img)
}

type Version int8

func (v Version) PatternSize() int {
	return 4*int(v) + 17
}

func LogError(err error) {
	log.Fatal(err)
}
