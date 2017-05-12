package compress

import (
	"errors"
	"fmt"
	"path/filepath"
)

var (
	ErrNotSupport = errors.New("Not suitable driver.")
)

var (
	// drivers is supported driver
	drivers []driver
)

type compressFunc func(dst, src string) error
type extractFunc func(dst, src string) error

type driver struct {
	format   string
	compress compressFunc
	extract  extractFunc
}

func RegisterFormat(format string, compress compressFunc, extract extractFunc) {
	fmt.Println("RegisterFormat: ", format)

	drivers = append(drivers, driver{format, compress, extract})
}

// get file format by file ext
func getFileFormat(filename string) string {
	return filepath.Ext(filename)
}

// match right driver with file ext
func matchDriver(format string) driver {
	fmt.Println("All drivers: ", drivers)

	for _, d := range drivers {
		if d.format == format {
			return d
		}
	}

	return driver{}
}

func Compress(dst, src string) error {
	ff := getFileFormat(dst)
	dvr := matchDriver(ff)

	if dvr.format == "" {
		return ErrNotSupport
	}

	return dvr.compress(dst, src)
}
