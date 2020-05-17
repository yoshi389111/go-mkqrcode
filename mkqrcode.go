package main

import (
	"image/png"
	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	qrCode, err := qr.Encode("Hello World", qr.M, qr.Auto)
	if err != nil {
		log.Fatal(err)
	}

	qrCode, err = barcode.Scale(qrCode, 200, 200)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("qrcode.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = png.Encode(file, qrCode)
	if err != nil {
		log.Fatal(err)
	}
}
