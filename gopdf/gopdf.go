package gopdf

import (
	"github.com/signintech/gopdf"
)

func Call() error {

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	// フォントの設定
	if err := setFont(&pdf); err != nil {
		return err
	}

	// ヘッダー・フッター
	setHeader(&pdf)
	setFooter(&pdf)

	// 内容
	content(&pdf)

	return pdf.WritePdf("./output/gopdf.pdf")
}

func content(pdf *gopdf.GoPdf) {
	pdf.AddPage()

	pdf.Cell(nil, "こんにちは")
}

func setFont(pdf *gopdf.GoPdf) error {
	err := pdf.AddTTFFont("notosans", "./font/NotoSansJP-VariableFont_wght.ttf")
	if err != nil {
		return err
	}

	err = pdf.SetFont("notosans", "", 14)
	if err != nil {
		return err
	}
	return nil
}

func setHeader(pdf *gopdf.GoPdf) {
	pdf.AddHeader(func() {
		pdf.SetY(5)
		pdf.Cell(nil, "header")
	})
}

func setFooter(pdf *gopdf.GoPdf) {
	pdf.AddFooter(func() {
		pdf.SetY(825)
		pdf.Cell(nil, "footer")
	})
}
