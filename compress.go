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

	// Optional parameter to specify the number of colors (2-256) for the GIF
	//
	// Specifying less colors compresses the GIF more.
	// If this is 0, gifsicle will default to using the image's color map.
	NumColors uint
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
		if o.NumColors >= 2 {
			gifsicleCli.NumColors(o.NumColors)
		}
	}

	return gifsicleCli.InputGif(g).Output(w).Run()
}
