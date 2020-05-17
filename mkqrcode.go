package main

import (
	"flag"
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	outFile := flag.String("o", "", "output file name")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	message := args[0]

	qrCode, err := qr.Encode(message, qr.M, qr.Auto)
	if err != nil {
		log.Fatal(err)
	}

	if *outFile != "" {
		err = outputQrCode(qrCode, 200, *outFile)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		black := "\033[40m  \033[0m"
		white := "\033[47m  \033[0m"
		printQrCode(qrCode, black, white)
	}
}

func outputQrCode(qrCode barcode.Barcode, size int, fileName string) error {
	qrCode, err := barcode.Scale(qrCode, size, size)
	if err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, qrCode)
	return err
}

func printQrCode(qrCode barcode.Barcode, black, white string) {
	rect := qrCode.Bounds()
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			if qrCode.At(x, y) == color.Black {
				fmt.Print(black)
			} else {
				fmt.Print(white)
			}
		}
		fmt.Println()
	}
}
