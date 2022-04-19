package main

import (
	"image/png"
	"net/http"
	"text/template"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type Page struct {
	Title string
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "QR CODE Generator"}

	t, _ := template.ParseFiles("generator.html")
	t.Execute(w, p)
}

func CodePage(w http.ResponseWriter, r *http.Request) {
	dataString := r.FormValue("dataString")
	qrCode, _ := qr.Encode(dataString, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)
	png.Encode(w, qrCode)
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/generator/", CodePage)
	http.ListenAndServe(":8088", nil)
}
