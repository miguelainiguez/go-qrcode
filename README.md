go-qrcode
=========

A (very) light golang convenience wrapper around [zbar](http://zbar.sourceforge.net/) made for shezadkhan137, used for qr code processing with rectangle coordinates.

## Requirements

To compile this package requires the zbar header files which can be installed on debian/ubuntu with
```
sudo apt-get install libzbar-dev
```

Go get the library:
```
go get github.com/miguelainiguez/go-qrcode/qrcode
```

## Usage (Currently under development)

It currently only supports extracting data from a PNG Image. Example Usage:

```go
import (
    "fmt"
    "github.com/miguelainiguez/go-qrcode/qrcode"
)

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
}
```

## Building

Building a staticlly linked binary with cgo dependencies can be a little fragile, by default cgo libraries are dynamically linked so require the libzbar-dev to be present on machine running your binary. However the following command may work if you want to statically link the zbar libs into your go binary.
```
go build -ldflags "-linkmode external -extldflags -static"
```
## How to run

go run main.go -i=test_images/boardTest9.png or if you made a build ./qrcode -i=test_images/boardTest9.png

## TODO

+ Add support for extrating qr data from video via V4L2
+ Add support for other image types
