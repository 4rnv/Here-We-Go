package main

import (
	"flag"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"strconv"
)

func drawCircle(img *image.RGBA, centerX, centerY, radius int, col color.Color) {
	for y := -radius; y <= radius; y++ {
		for x := -radius; x <= radius; x++ {
			if x*x+y*y <= radius*radius {
				img.Set(centerX+x, centerY+y, col)
			}
		}
	}
}

func hexToRGBA(hex string) color.RGBA {
	if hex[0] == '#' {
		hex = hex[1:]
	}
	if len(hex) != 6 {
		panic("Hex code must be 6 characters.")
	}
	r, err1 := strconv.ParseUint(hex[0:2], 16, 8)
	g, err2 := strconv.ParseUint(hex[2:4], 16, 8)
	b, err3 := strconv.ParseUint(hex[4:6], 16, 8)
	if err1 != nil || err2 != nil || err3 != nil {
		panic("Invalid hex code.")
	}
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}

func main() {
	// Pass -primary (hexcode), -secondary (hexcode), -filename (png) as CLI arguments
	const size = 1024
	const outerRadius = size / 2
	const innerRadius = size / 4
	primary := flag.String("primary", "", "primary color, inner circle")
	secondary := flag.String("secondary", "", "secondary color, outer circle")
	filename := flag.String("filename", "", "name of the image")
	flag.Parse()

	img := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Transparent}, image.Point{}, draw.Src)
	cx, cy := size/2, size/2
	drawCircle(img, cx, cy, outerRadius, hexToRGBA(*secondary))
	drawCircle(img, cx, cy, innerRadius, hexToRGBA(*primary))
	file, err := os.Create(*filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}
