package gofpdf

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func Call() error {
	// PDFを新規作成
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// フォントの設定
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello World")

	// ファイルに保存
	err := pdf.OutputFileAndClose("./output/gofpdf.pdf")
	if err != nil {
		return err
	}

	fmt.Println("PDFファイルが作成されました。")
	return nil
}
