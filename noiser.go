package main
import (
	"os"
	"flag"
	"image"
	"time"
    "image/png"
	"image/color"
	"math/rand"
)
var (
	width = flag.Int("w", 2480, "width")
	height = flag.Int("h", 3508, "height")
	randAlpha = flag.Int("a", 30, "rand alpha(MAX:1000)")
	outputName = flag.String("o", "texture.png", "output file name")
)
func fillRandAlpha(img *image.RGBA) {
    rect := img.Rect
    for y := rect.Min.Y; y < rect.Max.Y; y++ {
        for x := rect.Min.X; x < rect.Max.X; x++ {
			a := (uint8)(rand.Intn(255))
			if rand.Intn(1000) < 1000 - *randAlpha {
				a = 0
			}
            img.Set(x, y, color.RGBA{0, 255, 0, a})
        }
    }
}
func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
    img := image.NewRGBA(image.Rect(0, 0, *width, *height))
	fillRandAlpha(img)
	file, err := os.Create(*outputName)
	if err != nil {
		panic(err)
	}
    defer file.Close()
    if err := png.Encode(file, img); err != nil {
        panic(err)
    }
}