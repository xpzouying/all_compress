package tar

import (
	"fmt"
	"os"
	"testing"

	"github.com/xpzouying/compress"
)

func TestTarAFile(t *testing.T) {
	src := "tar_test.go"
	dst := "tar_test.go.tar"
	err := compress.Compress(dst, src)
	if err != nil {
		fmt.Println("Error: ", err)
		t.Error(err)
	} else {
		if err := os.Remove(dst); err != nil {
			panic(err)
		}
	}
}
