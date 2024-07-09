package benchmarks_test

import (
	"image/gif"
	"os"
	"path"
	"testing"
)

func BenchmarkGifDecodeAll(b *testing.B) {
	testFilePath := path.Join("..", "testfiles", "portrait_3mb.gif")

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			testFile, err := os.Open(testFilePath)
			if err != nil {
				b.Fatalf("os.Open: %v\n", err)
			}
			defer testFile.Close()

			_, err = gif.DecodeAll(testFile)
			if err != nil {
				b.Fatalf("gif.DecodeAll: %v\n", err)
			}
		}
	})
}
