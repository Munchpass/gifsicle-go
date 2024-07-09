package gifsicle

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
	gifsicleCli, err := NewGifsicle()
	if err != nil {
		return fmt.Errorf("NewGifsicle failed: %v", err)
	}

	if o != nil {
		gifsicleCli.Lossy(o.Lossy)
		gifsicleCli.OptimizeLevel(o.OptimizeLevel)
	}

	return gifsicleCli.InputGif(g).Output(w).Run()
}
