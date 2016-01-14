package main


import (
    "golang.org/x/tour/pic"
    "image/color"
    "image"
)
type Image struct{
	width, height int
    colr uint8
    
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func(i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) At(x, y int) color.Color {
    return color.RGBA{i.colr+uint8(x), i.colr+uint8(y), 255, 255}
}

func main() {
	m := Image{100, 100, 128}
	pic.ShowImage(m)
}
