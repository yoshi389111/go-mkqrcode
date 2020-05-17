package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/pborman/getopt/v2"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

const versionInfo = "@(#) $Id: mkqrcode.go 0.6.0 2020-05-17 15:34 yoshi389111 Exp $"

func main() {
	getopt.SetParameters("MESSAGE")
	outFile := getopt.StringLong("output", 'o', "", "output file name", "FILE")
	margin := getopt.IntLong("margin", 'm', 4, "margin of QR-code", "NUMBER")
	black := getopt.StringLong("black", 'b', "", "pattern of black", "STRING")
	white := getopt.StringLong("white", 'w', "", "pattern of white", "STRING")
	size := getopt.IntLong("size", 's', 200, "size of QR-code(Image)", "NUMBER")
	optLevel := getopt.EnumLong("level", 'l',
		[]string{"L", "M", "Q", "H"},
		"M", "error correction level", "{L|M|Q|H}")
	optEncoding := getopt.EnumLong("encoding", 'e',
		[]string{"auto", "numeric", "alphanumeric", "unicode"},
		"auto", "encoding of QR-code", "{auto|numeric|alphanumeric|unicode}")
	version := getopt.BoolLong("version", 'v', "show version info")
	help := getopt.BoolLong("help", 'h', "show help message")
	getopt.Parse()
	args := getopt.Args()
	if *help {
		getopt.PrintUsage(os.Stdout)
		os.Exit(0)
	}
	if *version {
		fmt.Println(versionInfo)
		os.Exit(0)
	}
	if len(args) != 1 {
		getopt.Usage()
		os.Exit(1)
	}
	if *margin < 0 {
		fmt.Fprintln(os.Stderr, "margin is negative")
		os.Exit(1)
	}
	if *black == "" && *white == "" {
		*black = "\033[40m  \033[0m"
		*white = "\033[47m  \033[0m"
	} else if *black == "" || *white == "" {
		fmt.Fprintln(os.Stderr, "specify both black and white")
		os.Exit(1)
	}
	var level qr.ErrorCorrectionLevel
	var encoding qr.Encoding

	switch strings.ToUpper(*optLevel) {
	case "L":
		level = qr.L
	case "M":
		level = qr.M
	case "Q":
		level = qr.Q
	case "H":
		level = qr.H
	default:
		fmt.Fprintf(os.Stderr, "invalid error correction level. (%s)\n", *optLevel)
		os.Exit(1)
	}
	switch strings.ToLower(*optEncoding) {
	case "auto":
		encoding = qr.Auto
	case "numeric":
		encoding = qr.Numeric
	case "alphanumeric":
		encoding = qr.AlphaNumeric
	case "unicode":
		encoding = qr.Unicode
	default:
		fmt.Fprintf(os.Stderr, "invalid encoding. (%s)\n", *optEncoding)
		os.Exit(1)
	}

	message := args[0]

	qrCode, err := qr.Encode(message, level, encoding)
	if err != nil {
		log.Fatal(err)
	}

	if *outFile != "" {
		err = outputQrCode(qrCode, *margin, *size, *outFile)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		printQrCode(qrCode, *margin, *black, *white)
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
