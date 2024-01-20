package main

import (
	"go-pdf-example/gofpdf"
	"go-pdf-example/gopdf"
	"go-pdf-example/wkhtmltopdf"
	"log"
)

const (
	execType = "wkhtmltopdf"
)

func main() {
	switch execType {
	case "gofpdf":
		if err := gofpdf.Call(); err != nil {
			log.Println(err)
		}

	case "gopdf":
		if err := gopdf.Call(); err != nil {
			log.Println(err)
		}

	case "wkhtmltopdf":
		if err := wkhtmltopdf.Call(); err != nil {
			log.Println(err)
		}
	}
}
