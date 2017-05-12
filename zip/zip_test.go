package zip

import (
	"fmt"
	"testing"

	"github.com/xpzouying/compress"
)

func TestZipAFile(t *testing.T) {
	err := compress.Compress("zip_test.go.zip", "zip_test.go")
	if err != nil {
		fmt.Println("Error: ", err)
		t.Error(err)
	}
}
