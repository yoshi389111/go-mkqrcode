# go-mkqrcode

This program is a CLI tool to generate QR-Code.

## Installation

```
go get github.com/yoshi389111/go-mkqrcode
```

## usage

```
$ go-mkqrcode --help
Usage: go-mkqrcode [options] MESSAGE
Options:
 -b, --black=STRING
                    pattern of black
 -e, --encoding={auto|numeric|alphanumeric|unicode}
                    encoding of QR-code [auto]
 -h, --help         show help message
 -l, --level={L|M|Q|H}
                    error correction level [M]
 -m, --margin=NUMBER
                    margin of QR-code [4]
 -o, --output=FILE  output file name
 -s, --size=NUMBER  size of QR-code(Image) [200]
 -v, --version      show version info
 -w, --white=STRING
                    pattern of white
```

## demo

Output QR-Code image file.

```
go-mkqrcode -o hello.png "hello"
```

Print QR-Code to terminal.

```
go-mkqrcode "hello world"
```

![hello world](https://raw.githubusercontent.com/wiki/yoshi389111/go-mkqrcode/images/hello_world.gif)

## Notice

### `github.com/boombuler/barcode`

* Copyright (c) 2014 Florian Sundermann
* [MIT Lisence](https://github.com/boombuler/barcode/blob/master/LICENSE)
* GitHub [boombuler/barcode](https://github.com/boombuler/barcode)
* [Document](https://pkg.go.dev/github.com/boombuler/barcode)

### `github.com/pborman/getopt`

* Copyright (c) 2017 Google Inc. All rights reserved.
* [BSD 3-Clause](https://github.com/pborman/getopt/blob/master/LICENSE)
* GitHub [pborman/getopt](https://github.com/pborman/getopt)
* [Document](https://pkg.go.dev/github.com/pborman/getopt)
