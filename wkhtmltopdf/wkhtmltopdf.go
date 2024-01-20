package wkhtmltopdf

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func Call() error {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}
	htmlSource := "./html/example.html"
	cssSource := "./html/example.css"

	page := wkhtmltopdf.NewPage(htmlSource)
	page.UserStyleSheet.Set(cssSource)
	pdfg.AddPage(page)

	// PDF作成
	err = pdfg.Create()
	if err != nil {
		return err
	}

	// 出力
	err = pdfg.WriteFile("./output/wkhtmltopdf.pdf")
	if err != nil {
		return err
	}

	return nil
}
