package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	outFile := flag.String("o", "", "output file name")
	margin := flag.Int("m", 4, "margin of QR-code")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}
	if *margin < 0 {
		fmt.Fprintln(os.Stderr, "margin is negative")
		os.Exit(1)
	}

	message := args[0]

	qrCode, err := qr.Encode(message, qr.M, qr.Auto)
	if err != nil {
		log.Fatal(err)
	}

	if *outFile != "" {
		err = outputQrCode(qrCode, *margin, 200, *outFile)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		black := "\033[40m  \033[0m"
		white := "\033[47m  \033[0m"
		printQrCode(qrCode, *margin, black, white)
	}
}

func outputQrCode(qrCode barcode.Barcode, margin, size int, fileName string) error {
	rect := qrCode.Bounds()
	cells := rect.Max.X + margin*2
	if size < cells {
		return fmt.Errorf("size is too small. required %d <= size", cells)
	}

	img := image.NewRGBA(image.Rect(0, 0, size, size))
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			cx := x*cells/size - margin
			cy := y*cells/size - margin
			if rect.Min.X <= cx && cx < rect.Max.X &&
				rect.Min.Y <= cy && cy < rect.Max.Y &&
				qrCode.At(cx, cy) == color.Black {

				img.Set(x, y, color.Black)
			} else {
				img.Set(x, y, color.White)
			}
		}
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	return err
}

func printQrCode(qrCode barcode.Barcode, margin int, black, white string) {
	rect := qrCode.Bounds()
	for y := rect.Min.Y - margin; y < rect.Max.Y+margin; y++ {
		for x := rect.Min.X - margin; x < rect.Max.X+margin; x++ {
			if rect.Min.X <= x && x < rect.Max.X &&
				rect.Min.Y <= y && y < rect.Max.Y &&
				qrCode.At(x, y) == color.Black {

				fmt.Print(black)
			} else {
				fmt.Print(white)
			}
		}
		fmt.Println()
	}
}
