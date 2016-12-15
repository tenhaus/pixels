package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func main() {

	file, _ := os.Open("test.jpg")
	image, err := jpeg.Decode(file)

	if err != nil {
		log.Fatal(err)
	}

	sortByReds(image)
}

func sortByReds(image image.Image) {
	size := image.Bounds().Size()
	xRange := make([]int, size.X)
	yRange := make([]int, size.Y)

	pixelIndex := make(map[uint32]color.Color)

	count := 0

	for xIndex := range xRange {
		for yIndex := range yRange {
			x := xIndex + 1
			y := yIndex + 1

			color := image.At(x, y)
			r, _, _, _ := color.RGBA()

			pixelIndex[r] = color
			count++
		}
	}

	fmt.Println(len(pixelIndex), count)

}

func write() {
}
