package qrcode

// #cgo LDFLAGS: -lzbar -lpng -ljpeg -lz -lrt -lm -pthread
// #include <stdio.h>
// #include <stdlib.h>
// #include <png.h>
// #include <zbar.h>
// #include "get_data.h"
// typedef void (*zbar_image_set_data_callback)(zbar_image_t *  image);
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type Result struct {
	Data       string
	Cr1        Coordinate
	Cr2        Coordinate
	Cl3        Coordinate
	Cl4        Coordinate
}

type Coordinate struct {
	X int
	Y int
}

func GetDataFromPNG(pngPath string) (results []Result, err error) {

	pth := C.CString(pngPath)
	scanner := C.zbar_image_scanner_create()
	C.zbar_image_scanner_set_config(scanner, 0, C.ZBAR_CFG_ENABLE, 1)

	defer C.zbar_image_scanner_destroy(scanner)

	var width, height C.int = 0, 0
	var raw unsafe.Pointer = nil
	errorCode := C.get_data(pth, &width, &height, &raw)
	if int(errorCode) != 0 {
		err = errors.New(fmt.Sprintf("Error reading from png file. Error code %d", errorCode))
		return
	}

	image := C.zbar_image_create()

	defer C.zbar_image_destroy(image)

	C.zbar_image_set_format(image, C.ulong(808466521))
	C.zbar_image_set_size(image, C.uint(width), C.uint(height))

	f := C.zbar_image_set_data_callback(C.zbar_image_free_data)
	C.zbar_image_set_data(image, raw, C.ulong(width*height), f)

	C.zbar_scan_image(scanner, image)

	symbol := C.zbar_image_first_symbol(image)
	for ; symbol != nil; symbol = C.zbar_symbol_next(symbol) {
		data := C.zbar_symbol_get_data(symbol)
		dataString := C.GoString(data)
		cr1 := Coordinate{int(C.zbar_symbol_get_loc_x(symbol, 0)), int(C.zbar_symbol_get_loc_y(symbol, 0))}
		cr2 := Coordinate{int(C.zbar_symbol_get_loc_x(symbol, 1)), int(C.zbar_symbol_get_loc_y(symbol, 1))}
		cl3 := Coordinate{int(C.zbar_symbol_get_loc_x(symbol, 2)), int(C.zbar_symbol_get_loc_y(symbol, 2))}
		cl4 := Coordinate{int(C.zbar_symbol_get_loc_x(symbol, 3)), int(C.zbar_symbol_get_loc_y(symbol, 3))}
		results = append(results, Result{dataString, cr1, cr2, cl3, cl4})
	}

	return
}
