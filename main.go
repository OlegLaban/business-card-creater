package main

import (
    "fmt"
    "log"

    "github.com/jung-kurt/gofpdf"
)

func main() {
    // Создаём PDF-документ
    pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(0, 0, 0)

    // Добавляем новую страницу
    pdf.AddPage()

    // Добавляем пример текста
    // pdf.MultiCell(0, 10, "", "", "", false)
	type Image struct {
        file    string
        offsetX float64
        offsetY float64
        width   float64
        height  float64
    }
    // Список изображений с их позициями и размерами
    images := make([]Image, 0)
	var yOffset float64 = 3
    var x, y float64 = 3, 3
	var width float64 = 40
	var count float64 = 0
    wd, hd, _ := pdf.PageSize(0)

	for {
		var columnCount float64 = 1
		for i := width; i <= float64(wd); i = i + width + x {
			images = append(images, Image{"images/image-Photoroom-with-border.png", x*columnCount + width * (columnCount - 1), y, width, 0})
			columnCount++
		}
		count++
		y = width * count + yOffset * (count + 1)
		fmt.Println(y, count, width)
		if (y + width + yOffset )  > hd {
			fmt.Println(y, count, width)
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
        pdf.ImageOptions(
            img.file,
            img.offsetX, img.offsetY,
            img.width, img.height,
            false,
            gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
            0,
            "", // Аннотация, оставлена пустой
        )
        
    }

    // Сохраняем PDF-документ
    err := pdf.OutputFileAndClose("output/output.pdf")
    if err != nil {
        log.Fatalf("Ошибка при создании PDF: %v", err)
    }

    fmt.Println("PDF успешно создан!")
}
