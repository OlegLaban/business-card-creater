package main

import (
	"log"
	"git.local/admin/image-to-pdf/internal/draw"
	"git.local/admin/image-to-pdf/pkg/img"
	"github.com/spf13/cobra"
)



func main() {
    orientation := "P"
    var uintStr string 
    pageMarginX := float64(3)
    pageMarginY := float64(3)
    pageOffsetX := float64(3)
    pageOffsetY := float64(3)
    pageFormat := "A4"
    var fontDir string
    var firstImagePath string
    var firstImageWidth float64
    var secondImagePath string
    var secondImageWidth float64
    var borderWidth int

    var rootCmd = &cobra.Command{
        Use: "main",
        Short: "Make pdf document with buisness card for printing",
        Run: func(cmd *cobra.Command, args[]string) {
            if borderWidth > 0 {
                var newFirstImgPath, newSecondImgPath string
                border := img.New(borderWidth, img.Black)

                newFirstImgPath, err := border.AddByPath(firstImagePath)
                if err != nil {
                    log.Fatalln("cannot add border for file " + firstImagePath)
                }

                firstImagePath = newFirstImgPath
                firstImageWidth = firstImageWidth + float64(2) * float64(borderWidth)

                if (secondImagePath != "") {
                    newSecondImgPath, err = border.AddByPath(secondImagePath)
                    if err != nil {
                        log.Fatalln("cannot add border for file " + secondImagePath)

                    }

                    secondImagePath = newSecondImgPath
                    secondImageWidth = secondImageWidth + float64(2) * float64(borderWidth)
                }
            }

            firstImage := &draw.ImageSettings{Filepath: firstImagePath, Width: firstImageWidth}
            var secondImage *draw.ImageSettings
            if secondImagePath != "" {
                if secondImageWidth == 0 {
                    secondImageWidth = firstImageWidth
                }
                secondImage = &draw.ImageSettings{Filepath: secondImagePath, Width: secondImageWidth}
            }

        
            pdfSettings := draw.PdfSettings{
                Orientation: orientation,
                Unit: uintStr,
                PageSettings: &draw.PageSettings{
                    OffsetX: pageOffsetX,
                    OffsetY: pageOffsetY,
                    MarginX: pageMarginX,
                    MarginY: pageMarginY,
                    Format: pageFormat,
                },
                FontDir: fontDir,
                FirstImage: firstImage,
                SecondImage: secondImage,
            }
            pdf := draw.New(pdfSettings)
            pdf.Draw()
        },
    }

    rootCmd.Flags().StringVarP(&firstImagePath, "fip", "f", "", "Path to first image")
    rootCmd.Flags().Float64VarP(&firstImageWidth, "fiw", "w", 0, "Width for image on page")
    
    rootCmd.Flags().StringVarP(&secondImagePath, "sip", "s", "", "Path to second image")
    rootCmd.Flags().Float64VarP(&secondImageWidth, "siw", "p", 0, "Width for second image on page")

    rootCmd.Flags().IntVarP(&borderWidth, "border", "b", 0, "Border width")
    
    rootCmd.Execute()
    
}