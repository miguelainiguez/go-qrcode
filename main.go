package main

import (
	"flag"
	"fmt"
	"github.com/miguelainiguez/go-qrcode/qrcode"
)

var image *string

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
		fmt.Printf("Data %s (%d,%d) \n", result.Data, result.Cl1.X, result.Cl1.Y )
	}
}
