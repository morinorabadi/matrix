package numberCreator

import (
	"image"
	"image/color"
)

func onColor(on bool, alpha uint8) color.RGBA {
	if on {
		return color.RGBA{
			R: 0,
			G: alpha,
			B: 0,
			A: 255,
		}
	} else {
		return color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 0,
		}
	}
}

func ZeroOrOne(number bool, alpha uint8) image.RGBA {
	width := 5
	height := 7

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	if number {
		img.Set(0, 0, onColor(false, alpha))
		img.Set(0, 1, onColor(false, alpha))
		img.Set(0, 2, onColor(false, alpha))
		img.Set(0, 3, onColor(false, alpha))
		img.Set(0, 4, onColor(false, alpha))
		img.Set(0, 5, onColor(false, alpha))
		img.Set(0, 6, onColor(false, alpha))

		img.Set(1, 0, onColor(true, alpha))
		img.Set(1, 1, onColor(false, alpha))
		img.Set(1, 2, onColor(false, alpha))
		img.Set(1, 3, onColor(false, alpha))
		img.Set(1, 4, onColor(false, alpha))
		img.Set(1, 5, onColor(false, alpha))
		img.Set(1, 6, onColor(true, alpha))

		img.Set(2, 0, onColor(true, alpha))
		img.Set(2, 1, onColor(true, alpha))
		img.Set(2, 2, onColor(true, alpha))
		img.Set(2, 3, onColor(true, alpha))
		img.Set(2, 4, onColor(true, alpha))
		img.Set(2, 5, onColor(true, alpha))
		img.Set(2, 6, onColor(true, alpha))

		img.Set(3, 0, onColor(false, alpha))
		img.Set(3, 1, onColor(false, alpha))
		img.Set(3, 2, onColor(false, alpha))
		img.Set(3, 3, onColor(false, alpha))
		img.Set(3, 4, onColor(false, alpha))
		img.Set(3, 5, onColor(false, alpha))
		img.Set(3, 6, onColor(true, alpha))

		img.Set(4, 0, onColor(false, alpha))
		img.Set(4, 1, onColor(false, alpha))
		img.Set(4, 2, onColor(false, alpha))
		img.Set(4, 3, onColor(false, alpha))
		img.Set(4, 4, onColor(false, alpha))
		img.Set(4, 5, onColor(false, alpha))
		img.Set(4, 6, onColor(false, alpha))
	} else {
		img.Set(0, 0, onColor(false, alpha))
		img.Set(0, 1, onColor(true, alpha))
		img.Set(0, 2, onColor(true, alpha))
		img.Set(0, 3, onColor(true, alpha))
		img.Set(0, 4, onColor(true, alpha))
		img.Set(0, 5, onColor(true, alpha))
		img.Set(0, 6, onColor(false, alpha))

		img.Set(1, 0, onColor(true, alpha))
		img.Set(1, 1, onColor(false, alpha))
		img.Set(1, 2, onColor(false, alpha))
		img.Set(1, 3, onColor(false, alpha))
		img.Set(1, 4, onColor(false, alpha))
		img.Set(1, 5, onColor(false, alpha))
		img.Set(1, 6, onColor(true, alpha))

		img.Set(2, 0, onColor(true, alpha))
		img.Set(2, 1, onColor(false, alpha))
		img.Set(2, 2, onColor(false, alpha))
		img.Set(2, 3, onColor(false, alpha))
		img.Set(2, 4, onColor(false, alpha))
		img.Set(2, 5, onColor(false, alpha))
		img.Set(2, 6, onColor(true, alpha))

		img.Set(3, 0, onColor(true, alpha))
		img.Set(3, 1, onColor(false, alpha))
		img.Set(3, 2, onColor(false, alpha))
		img.Set(3, 3, onColor(false, alpha))
		img.Set(3, 4, onColor(false, alpha))
		img.Set(3, 5, onColor(false, alpha))
		img.Set(3, 6, onColor(true, alpha))

		img.Set(4, 0, onColor(false, alpha))
		img.Set(4, 1, onColor(true, alpha))
		img.Set(4, 2, onColor(true, alpha))
		img.Set(4, 3, onColor(true, alpha))
		img.Set(4, 4, onColor(true, alpha))
		img.Set(4, 5, onColor(true, alpha))
		img.Set(4, 6, onColor(false, alpha))
	}

	return *img
}
