package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"sort"
)

type ColorSlice struct {
	original image.Image
	Colors   []color.Color
}

// Len is part of sort.Interface.
func (c *ColorSlice) Len() int {
	return len(c.Colors)
}

// Swap is part of sort.Interface.
func (c *ColorSlice) Swap(i, j int) {
	c.Colors[i], c.Colors[j] = c.Colors[j], c.Colors[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (c *ColorSlice) Less(i, j int) bool {
	_, r1, _, _ := c.Colors[i].RGBA()
	_, r2, _, _ := c.Colors[j].RGBA()

	return r1 < r2
}

func main() {

	file, _ := os.Open("test.jpg")
	image, err := jpeg.Decode(file)

	if err != nil {
		log.Fatal(err)
	}

	var colorSlice ColorSlice
	colorSlice.original = image
	buildSlice(image, &colorSlice)
	sort.Sort(&colorSlice)
	write(&colorSlice)
}

func buildSlice(image image.Image, colorSlice *ColorSlice) {
	size := image.Bounds().Size()
	xRange := make([]int, size.X)
	yRange := make([]int, size.Y)

	for xIndex := range xRange {
		for yIndex := range yRange {
			x := xIndex + 1
			y := yIndex + 1

			color := image.At(x, y)
			colorSlice.Colors = append(colorSlice.Colors, color)
		}
	}

}

func write(colors *ColorSlice) {
	newImage := image.NewRGBA(colors.original.Bounds())

	rowMax := colors.original.Bounds().Dx()

	rowCurrent := 1
	colCurrent := 1

	for _, currentColor := range colors.Colors {

		// Eh
		if rowCurrent > rowMax {
			rowCurrent = 1
			colCurrent++
		}

		newImage.Set(rowCurrent, colCurrent, currentColor)
		rowCurrent++
		// fmt.Println(rowCurrent, colCurrent, currentColor)
	}

	outfile, _ := os.Create("out.jpg")
	jpeg.Encode(outfile, newImage, &jpeg.Options{Quality: 100})

}
