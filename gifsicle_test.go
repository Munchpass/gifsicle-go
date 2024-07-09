package gifsicle_test

import (
	"image/gif"
	"os"
	"path"
	"testing"

	"github.com/Munchpass/gifsicle-go"
	"github.com/stretchr/testify/assert"
)

func validateGifBounds(t *testing.T, gifSource *gif.GIF, targetGifPath string) {
	fTarget, err := os.Open(targetGifPath)
	if !assert.Nil(t, err) {
		t.FailNow()
	}
	defer fTarget.Close()
	gifTarget, err := gif.DecodeAll(fTarget)
	if !assert.Nil(t, err) {
		t.FailNow()
	}

	assert.Equal(t, gifSource.Config.Height, gifTarget.Config.Height)
	assert.Equal(t, gifSource.Config.Width, gifTarget.Config.Width)
}

// Checks that the target gif can be parsed properly and has the same bounds
// as the source GIF.
func validateGifFiles(t *testing.T, sourceGifPath string, targetGifPath string) {
	//defer os.Remove("target.jpg")
	fSource, err := os.Open(sourceGifPath)
	if !assert.Nil(t, err) {
		t.FailNow()
	}
	gifSource, err := gif.DecodeAll(fSource)
	if !assert.Nil(t, err) {
		t.FailNow()
	}

	validateGifBounds(t, gifSource, targetGifPath)
}

func TestGifsicleVersion(t *testing.T) {
	g, err := gifsicle.NewGifsicle()
	if !assert.Nil(t, err) {
		t.FailNow()
	}
	v, err := g.Version()
	if !assert.Nil(t, err) {
		t.FailNow()
	}
	assert.NotEmpty(t, v)

	t.Logf("version: %s\n", v)
}

func TestGifsicleRunFromFile(t *testing.T) {
	g, err := gifsicle.NewGifsicle()
	if !assert.Nil(t, err) {
		t.FailNow()
	}

	sourceGif := path.Join("testfiles", "portrait_3mb.gif")
	outputGif := path.Join("testoutput", "portrait_3mb_output.gif")
	err = g.InputFile(sourceGif).
		OutputFile(outputGif).
		Lossy(80).
		OptimizeLevel(gifsicle.OPTIMIZE_LEVEL_THREE).
		Debug().
		Run()
	if !assert.Nil(t, err) {
		t.FailNow()
	}

	validateGifFiles(t, sourceGif, outputGif)
}
