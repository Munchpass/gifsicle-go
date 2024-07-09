package gifsiclego_test

import (
	"bytes"
	"image/gif"
	"io"
	"os"
	"path"
	"testing"

	gifsiclego "github.com/Munchpass/gifsicle-go"
)

func TestCompress(t *testing.T) {
	testFile, err := os.Open(path.Join("testfiles", "portrait_3mb.gif"))
	if err != nil {
		t.Fatalf("os.Open: %v\n", err)
	}

	decodedGif, err := gif.DecodeAll(testFile)
	if err != nil {
		t.Fatalf("gif.DecodeAll: %v\n", err)
	}

	var buf bytes.Buffer
	err = gifsiclego.Compress(&buf, decodedGif, &gifsiclego.Options{
		Lossy:         80,
		OptimizeLevel: gifsiclego.OPTIMIZE_LEVEL_THREE,
	})
	if err != nil {
		t.Fatalf("gifsiclego.Compress: %v\n", err)
	}

	// Need to re-read to get the original size
	testFile, err = os.Open(path.Join("testfiles", "portrait_3mb.gif"))
	if err != nil {
		t.Fatalf("os.Open: %v\n", err)
	}

	rawSource, err := io.ReadAll(testFile)
	if err != nil {
		t.Fatalf("io.ReadAll: %v\n", err)
	}

	t.Logf("size before: %d kb, size after: %d kb\n",
		len(rawSource)/1024, buf.Len()/1024)
}
