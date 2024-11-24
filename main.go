package main

import (
	"image"
	"image/color"
	"image/gif"
	"math/rand"
	"os"

	rope "github.com/morinorabadi/matrix/rope"
)

func main() {
	// Define image dimensions
	const width = 256
	const height = 256

	const totalFrames = 60

	// Slice to hold the frames
	var frames []*image.Paletted
	var delays []int

	// generate blinks
	totalBlinks := 800
	var Blinks []rope.Blink = make([]rope.Blink, totalBlinks)
	const blinkFrameLength = 5
	for i := 0; i < totalBlinks; i++ {
		startFrame := rand.Intn(totalFrames - blinkFrameLength)
		endFrame := startFrame + blinkFrameLength
		Blinks[i] = rope.GenerateBlink(startFrame, endFrame, rand.Intn(width), rand.Intn(width))
	}

	// Generate frames
	for frame := 0; frame < totalFrames; frame++ { // 10 frames

		plate := color.Palette{
			color.RGBA{R: 255, G: 0, B: 0},
			color.RGBA{R: 0, G: 0, B: 0},
			color.RGBA{R: 255, G: 255, B: 255},
		}
		for j := 0; j < 255; j += 64 {
			plate = append(plate, color.RGBA{R: 0, G: uint8(j), B: 0})
		}

		// Convert RGBA to Paletted
		paletted := image.NewPaletted(image.Rectangle{Min: image.Pt(0, 0), Max: image.Pt(255, 255)}, plate)
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				paletted.Set(x, y, color.Black)
			}
		}

		for _, Blink := range Blinks {
			Blink.Draw(paletted, frame)
		}

		frames = append(frames, paletted)
		delays = append(delays, 10) // 10 = 100ms delay per frame
	}

	// Create the output file
	outFile, err := os.Create("output.gif")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// Encode the GIF
	err = gif.EncodeAll(outFile, &gif.GIF{
		Image: frames,
		Delay: delays,
	})
	if err != nil {
		panic(err)
	}

	println("GIF saved as output.gif")
}
