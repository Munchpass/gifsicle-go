package benchmarks_test

import (
	"bytes"
	"image/gif"
	"os"
	"path"
	"testing"

	"github.com/brunoga/deep"
	"github.com/munchpass/gifsicle-go"
)

func BenchmarkCompress(b *testing.B) {
	testFilePath := path.Join("..", "testfiles", "portrait_3mb.gif")
	testFile, err := os.Open(testFilePath)
	if err != nil {
		b.Fatalf("os.Open: %v\n", err)
	}
	defer testFile.Close()

	decodedGif, err := gif.DecodeAll(testFile)
	if err != nil {
		b.Fatalf("gif.DecodeAll: %v\n", err)
	}

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			var buf bytes.Buffer
			err = gifsicle.Compress(&buf, decodedGif, &gifsicle.Options{
				Lossy:         80,
				OptimizeLevel: gifsicle.OPTIMIZE_LEVEL_THREE,
			})
			if err != nil {
				b.Fatalf("gifsicle.Compress: %v\n", err)
			}
		}
	})
}

func BenchmarkCompressFromReader(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			testFilePath := path.Join("..", "testfiles", "portrait_3mb.gif")
			testFile, err := os.Open(testFilePath)
			if err != nil {
				b.Fatalf("os.Open: %v\n", err)
			}
			defer testFile.Close()

			var buf bytes.Buffer
			err = gifsicle.CompressFromReader(&buf, testFile, &gifsicle.Options{
				Lossy:         80,
				OptimizeLevel: gifsicle.OPTIMIZE_LEVEL_THREE,
			})
			if err != nil {
				b.Fatalf("gifsicle.Compress: %v\n", err)
			}
		}
	})
}

func BenchmarkReuseGifsicleCompress(b *testing.B) {
	gifsicleCli, err := gifsicle.NewGifsicle()
	if err != nil {
		b.Fatalf("gifsicle.NewGifsicle: %v\n", err)
	}

	gifsicleCli = gifsicleCli.Lossy(80).OptimizeLevel(gifsicle.OPTIMIZE_LEVEL_THREE)
	testFilePath := path.Join("..", "testfiles", "portrait_3mb.gif")
	testFile, err := os.Open(testFilePath)
	if err != nil {
		b.Fatalf("os.Open: %v\n", err)
	}
	defer testFile.Close()

	decodedGif, err := gif.DecodeAll(testFile)
	if err != nil {
		b.Fatalf("gif.DecodeAll: %v\n", err)
	}

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			// Necessary because normal gifsicleCli.Run is not thread-safe
			cli, err := deep.Copy(gifsicleCli)
			if err != nil {
				b.Fatalf("deep.Copy failed: %v\n", err)
			}

			var buf bytes.Buffer
			err = cli.InputGif(decodedGif).Output(&buf).Run()
			if err != nil {
				b.Fatalf("gifsicleCli.Run: %v\n", err)
			}
		}
	})
}
