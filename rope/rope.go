package rope

import (
	"image"
	"math/rand"
)

type Rope struct {
	startFrame int
	blinks     []Blink
}

func (r Rope) Draw(paletted *image.Paletted, frame int) {
	for _, blink := range r.blinks {
		blink.Draw(paletted, frame)
	}
}

func GenerateRope(startPosX int, startPosY int, startFrame int, totalFrames int) Rope {
	// generate blinks
	totalBlinks := 25 + rand.Intn(50)
	var Blinks []Blink = make([]Blink, totalBlinks)
	for i := 0; i < totalBlinks; i++ {
		startFrame := startFrame + i
		Blinks[i] = GenerateBlink(startFrame, startFrame+10, startPosX, startPosY+i*9)
	}

	return Rope{startFrame: startFrame, blinks: Blinks}
}
