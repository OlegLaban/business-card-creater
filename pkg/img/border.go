package img

import (
	"fmt"
	"image"
	icolor "image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	Black = "Black"
)

type Border struct {
	width int
	color string
}

type color struct {
	r uint8
	g uint8
	b uint8
}

func New(width int, color string) *Border {
	return &Border{width: width, color: color}
}

func (b *Border) AddByPath(path string) (string, error) {
	var newPath string

	r, err := os.Open(path)

	if err != nil {
		return "", err
	}
	defer r.Close()

	newPath, err = b.Add(r)
	if err != nil {
		log.Default().Println(err)
		return "", err
	}

	return newPath, nil
}

func (b *Border) Add(file io.Reader) (string, error) {
	srcImage, _, err := image.Decode(file)
	
	if err != nil {
		return "", err
	}

	bounds := srcImage.Bounds()
	newWidth := bounds.Dx() + 2*b.width
	newHeight := bounds.Dy() + 2*b.width

	destImage := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	needColor := b.getColor()

	borderColor := icolor.RGBA{needColor.r, needColor.g, needColor.g, 255}
	// Заливаем границу по краям
	// Верхняя граница
	draw.Draw(destImage, image.Rect(0, 0, newWidth, b.width), &image.Uniform{borderColor}, image.Point{}, draw.Src)
	// Нижняя граница
	draw.Draw(destImage, image.Rect(0, newHeight-b.width, newWidth, newHeight), &image.Uniform{borderColor}, image.Point{}, draw.Src)
	// Левая граница
	draw.Draw(destImage, image.Rect(0, 0, b.width, newHeight), &image.Uniform{borderColor}, image.Point{}, draw.Src)
	// Правая граница
	draw.Draw(destImage, image.Rect(newWidth-b.width, 0, newWidth, newHeight), &image.Uniform{borderColor}, image.Point{}, draw.Src)

	// Вставляем исходное изображение в центр нового изображения
	offset := image.Pt(b.width, b.width)
	draw.Draw(destImage, bounds.Add(offset), srcImage, bounds.Min, draw.Over)

	outputFile, err := os.CreateTemp("", strconv.Itoa(int(time.Now().Unix()))+".png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, destImage)

	if err != nil {
		return "", err
	}

	fmt.Println("Граница добавлена")

	return outputFile.Name(), nil
}

func (b *Border) getColor() *color {
	var c *color

	switch b.color {
	case "Black":
		c = &color{r: 0, g: 0, b: 0}
	}

	return c
}
