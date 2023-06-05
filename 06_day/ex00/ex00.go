package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	const width, hight = 300, 300

	img := image.NewNRGBA(image.Rect(0, 0, width, hight))

	for y := 0; y < hight; y++ {
		for x := 0; x < width; x++ {
			if x < 10 || y < 10 || x > 290 || y > 290 {
				img.Set(x, y, color.Black)
			} else if y < 40 || x < 40 || x > 260 || y > 260 {
				img.Set(x, y, color.Gray{Y: uint8((x) << 3 & 255)})
			} else if y >= 40 && y < 50 || x >= 40 && x < 50 || x > 250 && x <= 260 || y > 250 && y <= 260 {
				img.Set(x, y, color.White)
			} else {
				img.Set(x, y, color.NRGBA{
					R: uint8((y) & 255),
					G: uint8((x) << 1 & 255),
					B: uint8((y) << 1 & 255),
					A: 255,
				})
			}
		}
	}
	f, err := os.Create("amazing_logo.png")
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
