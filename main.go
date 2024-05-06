package main

import (
	"go-unidoc/configs"
	"os"
	"path/filepath"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/model"
)

func init() {
	err := license.SetMeteredKey(configs.UNIDOC_API_KEY)
	if err != nil {
		panic(err)
	}
}
func main() {
	basePath := "./pdf/"
	sampleFiles := []string{"s1.pdf", "s2.pdf", "s3.pdf", "s4.pdf"}
	// Merge the sample files
	outputPath := "./out/merged.pdf"
	newInput := []string{}
	for _, file := range sampleFiles {
		newInput = append(newInput, filepath.Join(basePath, file))
	}

	if err := mergePdf(newInput, outputPath); err != nil {
		panic("mergePdf failed: " + err.Error())
	}

}

// https://docs.unidoc.io/docs/unipdf/guides/page-manipulation/pdf_merge/
func mergePdf(inputPaths []string, outputPath string) error {
	pdfWriter := model.NewPdfWriter()

	for _, inputPath := range inputPaths {
		pdfReader, f, err := model.NewPdfReaderFromFile(inputPath, nil)
		if err != nil {
			return err
		}
		defer f.Close()

		numPages, err := pdfReader.GetNumPages()
		if err != nil {
			return err
		}

		for i := 0; i < numPages; i++ {
			pageNum := i + 1

			page, err := pdfReader.GetPage(pageNum)
			if err != nil {
				return err
			}

			err = pdfWriter.AddPage(page)
			if err != nil {
				return err
			}
		}
	}

	fWrite, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	defer fWrite.Close()

	err = pdfWriter.Write(fWrite)
	if err != nil {
		return err
	}

	return nil
}
