package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func mergeImagesToPDF(images []string, outputPath string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	for _, imagePath := range images {
		pdf.AddPage()
		ext := strings.ToLower(filepath.Ext(imagePath))
		imageType := "JPG"
		if ext == ".png" {
			imageType = "PNG"
		}
		pdf.ImageOptions(imagePath, 10, 10, 190, 0, false, gofpdf.ImageOptions{ImageType: imageType}, 0, "")
	}
	return pdf.OutputFileAndClose(outputPath)
}

func mergePDFs(inputPaths []string, outputPath string) error {
	return api.MergeCreateFile(inputPaths, outputPath, false, nil)
}

func getFilesFromFolder(folder string) (images []string, pdfs []string, err error) {
	files, err := os.ReadDir(folder)
	if err != nil {
		return nil, nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(file.Name()))
		fullPath := filepath.Join(folder, file.Name())
		switch ext {
		case ".jpg", ".jpeg", ".png":
			images = append(images, fullPath)
		case ".pdf":
			pdfs = append(pdfs, fullPath)
		}
	}

	return images, pdfs, nil
}

func main() {
	inputFolder := "./input"          // โฟลเดอร์ input
	tempImagePDF := "temp_images.pdf" // ไฟล์ชั่วคราวสำหรับรูปภาพ
	outputFinalPDF := "output.pdf"

	images, pdfs, err := getFilesFromFolder(inputFolder)
	if err != nil {
		log.Fatalf("Error reading input folder: %v\n", err)
	}

	// ถ้ามีรูปภาพ ให้รวมรูปภาพเป็น PDF ชั่วคราว
	if len(images) > 0 {
		if err := mergeImagesToPDF(images, tempImagePDF); err != nil {
			log.Fatalf("Error merging images to PDF: %v\n", err)
		}
		// เพิ่ม PDF ที่ได้จากรูปภาพไปยังรายการ PDF
		pdfs = append([]string{tempImagePDF}, pdfs...)
	}

	// รวม PDF ทั้งหมด
	if len(pdfs) > 0 {
		if err := mergePDFs(pdfs, outputFinalPDF); err != nil {
			log.Fatalf("Error merging PDFs: %v\n", err)
		}
		fmt.Println("Merged PDF created successfully at:", outputFinalPDF)
	} else {
		fmt.Println("No PDFs to merge.")
	}

	// ลบไฟล์ชั่วคราว
	if len(images) > 0 {
		os.Remove(tempImagePDF)
	}
}
