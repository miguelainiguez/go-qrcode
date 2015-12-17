package main

import (
	"flag"
	"fmt"
	"github.com/miguelainiguez/go-qrcode/qrcode"
)

var image *string

//go build -o test -ldflags "-linkmode external -extldflags -static"
func init() {
	image = flag.String("i", "", "image path")
}

func main() {

	flag.Parse()

	results, err := qrcode.GetDataFromPNG(*image)
	if err != nil {
		panic(err)
	}

	for _, result := range results{
		//Print Data and coordinate first corner right
		fmt.Printf("Data %s (%d,%d) \n", result.Data, result.Cr1.X, result.Cr1.Y )
	}
}
