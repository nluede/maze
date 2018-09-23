package imageprocession

import (
	"github.com/nluede/maze/maze"
	"image"
	_ "image/png"
	"os"
)

// ReadImage reads the data from an image and creates a maze for this data. The function is designed to work on a black
// and white image. However, every white pixel of the image will be evaluated as a passway and every other pixel will
// be evaluated as a wall. The alpha values are omitted.
func ReadImage(filePath string) (maze.Maze, error) {
	infile, err := os.Open(filePath)

	if err != nil {
		return maze.New(0, 0), err
	}
	defer infile.Close()

	src, _, err := image.Decode(infile)
	if err != nil {
		return maze.New(0, 0), err
	}

	bounds := src.Bounds()
	result := maze.New(bounds.Dx(), bounds.Dy())

	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			pixel := src.At(x, y)
			r, g, b, _ := pixel.RGBA()
			if r == 0xffff && g == 0xffff && b == 0xffff {
				result.Set(x, y, true)
			}
		}
	}

	return result, nil
}
