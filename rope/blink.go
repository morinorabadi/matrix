package rope

import (
	"image"
	"math/rand"

	numberCreator "github.com/morinorabadi/matrix/numberCreator"
)

type Blink struct {
	startFrame int
	endFrame   int
	length     int
	offsetX    int
	offsetY    int
}

func (b Blink) Draw(paletted *image.Paletted, frame int) {

	if frame > b.endFrame {
		return
	}

	if frame < b.startFrame {
		return
	}
	alpha := float32(frame-b.startFrame) / float32(b.length)
	img := numberCreator.ZeroOrOne(rand.Intn(10), uint8(alpha*255.0))
	for ix := 0; ix < 5; ix++ {
		for iy := 0; iy < 7; iy++ {
			paletted.Set(b.offsetX+ix, b.offsetY+iy, img.At(ix, iy))
		}
	}
}

func GenerateBlink(
	startFrame int,
	endFrame int,
	offsetX int,
	offsetY int,
) Blink {
	blink := Blink{
		startFrame: startFrame,
		endFrame:   endFrame,
		offsetX:    offsetX,
		offsetY:    offsetY,
		length:     endFrame - startFrame,
	}
	return blink
}
