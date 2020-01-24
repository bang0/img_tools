package main
import (
	"os"
	"flag"
	"image"
    "image/png"
	"image/color"
)
var(
	outputName = flag.String("o", "dst.png", "output file name")
	alphaBorder = flag.Int("a", 0, "alpha border")
)
func main() {
    flag.Parse()
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
    if err != nil {
        panic(err)
	}
	rect := img.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, rect.Max.X, rect.Max.Y))
    for y := rect.Min.Y; y < rect.Max.Y; y++ {
        for x := rect.Min.X; x < rect.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			if a > (uint32)(*alphaBorder) {
				dst.Set(x, y, color.RGBA{(uint8)(r), (uint8)(g), (uint8)(b), 255})
			}
        }
	}
	dstfile, err := os.Create(*outputName)
	if err != nil {
		panic(err)
	}
    defer dstfile.Close()
    if err := png.Encode(dstfile, dst); err != nil {
        panic(err)
    }
}