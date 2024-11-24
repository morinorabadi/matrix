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
	// Define gif dimensions
	/*
		change this properties to generate your rain matrix
	*/
	const width = 600
	const height = 400
	const totalFrames = 30

	// Slice to hold the frames
	var frames []*image.Paletted
	var delays []int

	totalRopes := width * height * totalFrames / 40000
	var ropes []rope.Rope = make([]rope.Rope, totalRopes)
	for i := 0; i < totalRopes; i++ {
		startFrame := rand.Intn(totalFrames+60) - 30
		rWidth := rand.Intn(width/9) * 9
		rHeight := rand.Intn(height/9) * 9
		rHeight += rand.Intn(50) - 50
		ropes[i] = rope.GenerateRope(rWidth, rHeight, startFrame, totalFrames)
	}

	plate := color.Palette{
		color.RGBA{R: 255, G: 0, B: 0},
		color.RGBA{R: 0, G: 0, B: 0},
		color.RGBA{R: 255, G: 255, B: 255},
	}
	for j := 0; j <= 255; j += 16 {
		plate = append(plate, color.RGBA{R: 0, G: uint8(j), B: 0})
	}

	// Generate frames
	for frame := 0; frame < totalFrames; frame++ { // 10 frames

		// Convert RGBA to Paletted
		paletted := image.NewPaletted(image.Rectangle{Min: image.Pt(0, 0), Max: image.Pt(width, height)}, plate)
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				paletted.Set(x, y, color.Black)
			}
		}

		for _, Blink := range ropes {
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
