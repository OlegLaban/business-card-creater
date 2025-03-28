package draw

import (
	"fmt"
	"log"

	"github.com/jung-kurt/gofpdf"
)

type Image struct {
    file    string
    offsetX float64
    offsetY float64
    width   float64
    height  float64
}

type ImageSettings struct {
    Filepath string
    Width float64
}

type PageSettings struct {
    OffsetX float64
    OffsetY float64
    MarginX float64
    MarginY float64
	Format string
}

type PdfSettings struct {
	Orientation string
	Unit string
	PageSettings *PageSettings
	FontDir string
	FirstImage *ImageSettings
	SecondImage *ImageSettings
}

type PdfDraw struct {
	pdf *gofpdf.Fpdf
	settings PdfSettings
}

func New(settings PdfSettings) *PdfDraw {
	pdf := gofpdf.New(settings.Orientation, settings.Unit, settings.PageSettings.Format, settings.FontDir)

	return &PdfDraw{pdf: pdf, settings: settings}
}

func (pd *PdfDraw) Draw() {
    // Создаём PDF-документ
    
	pd.pdf.SetMargins(0, 0, 0)

    // Добавляем новую страницу
    pd.pdf.AddPage()

    pd.drawSide(pd.settings.FirstImage)

	if pd.settings.SecondImage != nil {
		pd.pdf.AddPage()

		wd, _, _ := pd.pdf.PageSize(pd.pdf.PageNo())
		pd.pdf.TransformBegin()
		pd.pdf.TransformMirrorHorizontal(wd/2)
		pd.drawSide(pd.settings.SecondImage)
		pd.pdf.TransformEnd()
	}
    
    // Сохраняем PDF-документ
    err := pd.pdf.OutputFileAndClose("output.pdf")
    if err != nil {
        log.Fatalf("Ошибка при создании PDF: %v", err)
    }

    fmt.Println("PDF успешно создан!")
}

func (pd *PdfDraw) drawSide(imgSettings *ImageSettings) {
    // Список изображений с их позициями и размерами
    images := make([]Image, 0)
    var y float64 = pd.settings.PageSettings.OffsetY
	var count float64 = 0
    wd, hd, _ := pd.pdf.PageSize(0)

	for {
		var columnCount float64 = 1
		for i := imgSettings.Width; i <= float64(wd); i = i + imgSettings.Width + pd.settings.PageSettings.OffsetX {
            image := Image{
                imgSettings.Filepath,
                pd.settings.PageSettings.OffsetX*columnCount + imgSettings.Width * (columnCount - 1),
                y,
                imgSettings.Width,
                0,
            }
			images = append(images, image)
			columnCount++
		}
        
		count++

		y = imgSettings.Width * count + pd.settings.PageSettings.OffsetX * (count + 1)
		
		if (y + imgSettings.Width + pd.settings.PageSettings.OffsetY)  > hd {
			break
		}

		
	}

    // Обрабатываем каждое изображение
    for _, img := range images {
        // // Проверяем, что файл существует
        // if pdf.ImageInfo(img.file) == nil {
        //     log.Fatalf("Файл изображения не найден: %s", img.file)
        // }

        // Добавляем изображение
        pd.pdf.ImageOptions(
            img.file,
            img.offsetX, img.offsetY,
            img.width, img.height,
            false,
            gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
            0,
            "", // Аннотация, оставлена пустой
        )
    }
}