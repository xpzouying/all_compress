// Implement of compress with zip.

package zip

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/xpzouying/compress"
)

var (
	ErrSrcNotExists     = errors.New("source file not exists")
	ErrDstExists        = errors.New("dest file is exists")
	ErrFileFormat       = errors.New("error file format")
	ErrSrcNotSupportDir = errors.New("not support dir yet")
)

// Compress file
func Compress(dst, src string) error {
	fmt.Println("Use zip to compress.")

	// make sure src file is exists
	srcInfo, err := os.Stat(src)
	if os.IsNotExist(err) {

		return ErrSrcNotExists
	}

	// make sure dest file is not exists
	if _, err := os.Stat(dst); os.IsExist(err) {
		fmt.Println(err)

		return ErrDstExists
	}

	// make sure dest file is *.zip
	ff := filepath.Ext(dst)
	if ".zip" != ff {
		return ErrFileFormat
	}

	// compress file
	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer d.Close()

	w := zip.NewWriter(d)
	defer w.Close()

	// create a file in `zip`.
	// create file header use src.
	fileInZip, err := w.Create(srcInfo.Name())
	if err != nil {
		return err
	}

	// compress src file
	if srcInfo.IsDir() {
		return ErrSrcNotSupportDir
	}

	// construct file
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	_, err = io.Copy(fileInZip, srcFile)
	if err != nil {
		return err
	}

	return nil
}

// Extract file
func Extract(dst, src string) error {
	return nil
}

// Return abs dir
// if dst has not dir path, then use current dir
func absDir(dst string) (string, error) {
	abspath, err := filepath.Abs(dst)
	if err != nil {
		return "", err
	}

	return filepath.Dir(abspath), nil
}

func init() {
	compress.RegisterFormat(".zip", Compress, Extract)
}
