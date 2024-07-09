package gifsicle_test

import (
	"bytes"
	"image/gif"
	"io"
	"os"
	"path"
	"testing"

	"github.com/Munchpass/gifsicle-go"
)

func TestCompress(t *testing.T) {
	testFilePath := path.Join("testfiles", "portrait_3mb.gif")
	testFile, err := os.Open(testFilePath)
	if err != nil {
		t.Fatalf("os.Open: %v\n", err)
	}

	decodedGif, err := gif.DecodeAll(testFile)
	if err != nil {
		t.Fatalf("gif.DecodeAll: %v\n", err)
	}

	var buf bytes.Buffer
	err = gifsicle.Compress(&buf, decodedGif, &gifsicle.Options{
		Lossy:         80,
		OptimizeLevel: gifsicle.OPTIMIZE_LEVEL_THREE,
	})
	if err != nil {
		t.Fatalf("gifsicle.Compress: %v\n", err)
	}

	// Need to re-read to get the original size
	testFile, err = os.Open(testFilePath)
	if err != nil {
		t.Fatalf("os.Open: %v\n", err)
	}

	rawSource, err := io.ReadAll(testFile)
	if err != nil {
		t.Fatalf("io.ReadAll: %v\n", err)
	}

	t.Logf("size before: %d kb, size after: %d kb\n",
		len(rawSource)/1024, buf.Len()/1024)

	testFilePath = path.Join("testfiles", "portrait_1dot5mb.gif")
	testFile, err = os.Open(testFilePath)
	if err != nil {
		t.Fatalf("os.Open: %v\n", err)
	}

	decodedGif, err = gif.DecodeAll(testFile)
	if err != nil {
		t.Fatalf("gif.DecodeAll: %v\n", err)
	}

	buf.Reset()
	err = gifsicle.Compress(&buf, decodedGif, &gifsicle.Options{
		Lossy:         120,
		OptimizeLevel: gifsicle.OPTIMIZE_LEVEL_THREE,
		NumColors:     256,
	})
	if err != nil {
		t.Fatalf("gifsicle.Compress: %v\n", err)
	}

	// Need to re-read to get the original size
	testFile, err = os.Open(testFilePath)
	if err != nil {
		t.Fatalf("os.Open: %v\n", err)
	}

	rawSource, err = io.ReadAll(testFile)
	if err != nil {
		t.Fatalf("io.ReadAll: %v\n", err)
	}

	t.Logf("size before: %d kb, size after: %d kb\n",
		len(rawSource)/1024, buf.Len()/1024)
}
