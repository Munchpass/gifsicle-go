package gifsiclego

import (
	"fmt"
	"image/gif"
	"io"
)

type Options struct {
	// This parameter ranges from 0-200 (higher value -> more compression).
	//
	// This is the main compression parameter used in websites like ezgif!
	Lossy uint

	// Additional compression optimization.
	OptimizeLevel OptimizeLevel
}

// Shortcut function to compress GIFs quickly and easily with gifsicle.
func Compress(w io.Writer, g *gif.GIF, o *Options) error {
	gifsicle, err := NewGifsicle()
	if err != nil {
		return fmt.Errorf("NewGifsicle failed: %v", err)
	}

	if o != nil {
		gifsicle.Lossy(o.Lossy)
		gifsicle.OptimizeLevel(o.OptimizeLevel)
	}

	return gifsicle.InputGif(g).Output(w).Run()
}
