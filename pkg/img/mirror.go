package img

import (
	"image"
	"image/png"
	"io"
	"os"
	"strconv"
	"time"
)

type Mirror struct {}

func NewMirror() *Mirror {
	return &Mirror{}
}

func mirrorImageH(img image.Image) *image.RGBA {
	bounds := img.Bounds()
    mirrored := image.NewRGBA(bounds)

    // Копируем пиксели зеркально
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            srcX := bounds.Max.X - 1 - x
            mirrored.Set(x, y, img.At(srcX, y))
        }
    }

    return mirrored
}

func (m Mirror) Mirror(file io.Reader) (string, error) {
	img, _, err := image.Decode(file)

	if  err != nil {
		return "", err
	}

	mirrorImag := mirrorImageH(img)

	outFile, err := os.CreateTemp("", strconv.Itoa(int(time.Now().Unix()))+".png")

	if err != nil {
		return "", err
	}
	defer outFile.Close()

	err = png.Encode(outFile, mirrorImag)
	if err != nil {
		return "", err
	}

	return outFile.Name(), nil
}