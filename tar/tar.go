package tar

import (
	"archive/tar"
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

func Compress(dst, src string) error {
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

	// compress src file
	if srcInfo.IsDir() {
		return ErrSrcNotSupportDir
	}

	// make sure dest file is *.zip
	ff := filepath.Ext(dst)
	if ".tar" != ff {
		return ErrFileFormat
	}

	// compress file
	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer d.Close()

	w := tar.NewWriter(d)
	defer w.Close()

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}

	// write tar file
	// NOTE: Must add header, else error: write too long
	hdr, err := tar.FileInfoHeader(srcInfo, "")
	if err != nil {
		return err
	}

	err = w.WriteHeader(hdr)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, srcFile)
	if err != nil {
		return err
	}

	return nil
}

// Extract file
func Extract(dst, src string) error {
	return nil
}

func init() {
	compress.RegisterFormat(".tar", Compress, Extract)
}
