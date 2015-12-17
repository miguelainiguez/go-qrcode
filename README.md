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
    results, err := qrcode.GetDataFromPNG("path/to/image.png")
    if err != nil {
        panic(err)
    }

    for _, result := range results{
        fmt.Printf("Symbol Type: %s, Data %s", result.SymbolType, result.Data )
    }

}
```

## Building

Building a staticlly linked binary with cgo dependencies can be a little fragile, by default cgo libraries are dynamically linked so require the libzbar-dev to be present on machine running your binary. However the following command may work if you want to statically link the zbar libs into your go binary.
```
go build -ldflags "-linkmode external -extldflags -static"
```

## TODO

+ Add support for extrating qr data from video via V4L2
+ Add support for other image types



