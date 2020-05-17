package main

import (
	"flag"
	"fmt"
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
	if *outFile == "" {
		fmt.Fprintln(os.Stderr, "output file is required.")
		os.Exit(1)
	}

	message := args[0]

	qrCode, err := qr.Encode(message, qr.M, qr.Auto)
	if err != nil {
		log.Fatal(err)
	}

	qrCode, err = barcode.Scale(qrCode, 200, 200)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(*outFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = png.Encode(file, qrCode)
	if err != nil {
		log.Fatal(err)
	}
}
