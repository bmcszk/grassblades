package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

func main() {
	// Generate a grass image and save it to a PNG file
	grassImg := generateGrass()
	grassFile, _ := os.Create("grass.png")
	defer grassFile.Close()
	png.Encode(grassFile, grassImg)

	// Generate a dirt image and save it to a PNG file
	dirtImg := generateDirt()
	dirtFile, _ := os.Create("dirt.png")
	defer dirtFile.Close()
	png.Encode(dirtFile, dirtImg)

}

func generateGrass() *image.RGBA {
	width, height := 32, 32
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill the image with a random shade of green
	green := color.RGBA{0, uint8(rand.Intn(200)), 0, 255}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, green)
		}
	}

	// Randomly generate grass blades
	for i := 0; i < 50; i++ {
		x := rand.Intn(width)
		y := rand.Intn(height-5) + 5
		length := rand.Intn(4) + 1
		for j := 0; j < length; j++ {
			img.Set(x, y-j, color.RGBA{0, 255, 0, 255})
		}
	}

	return img
}

func generateDirt() *image.RGBA {
	width, height := 32, 32
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill the image with a random shade of brown
	brown := color.RGBA{139, 69, 19, 255}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r := uint8(rand.Intn(40)) + brown.R
			g := uint8(rand.Intn(20)) + brown.G
			b := uint8(rand.Intn(20)) + brown.B
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	return img
}
