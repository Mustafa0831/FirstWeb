package main

import (
	AsciiArt "AsciiArt/asciiart"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var mainTempalate, errorTempalate *template.Template

func main() {
	var err error

	mainTempalate, err = template.ParseFiles("PageTemplate/index.html")
	if err != nil {
		fmt.Println("No template")
		return
	}

	errorTempalate, err = template.ParseFiles("PageTemplate/error.html")
	if err != nil {
		fmt.Println("No template")
		return
	}

	http.HandleFunc("/", IndexPage)
	http.HandleFunc("/ascii-art", ConvertText)
	fmt.Println("server listening on port 8001...")
	log.Fatal(http.ListenAndServe(":8001", nil))
}

//ConvertText is executing request(respond to request)
func ConvertText(w http.ResponseWriter, r *http.Request) {
	//checking method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		errorTempalate.Execute(w, "405: Status method not allowed")
		return
	}

	//Getting FontStyle
	font := r.FormValue("banner")
	if font != "standard" && font != "shadow" && font != "thinkertoy" {
		w.WriteHeader(http.StatusBadRequest)
		errorTempalate.Execute(w, "400: Bad Request")
		return
	}

	//Getting Text
	text := r.FormValue("data")
	if text == "" {
		w.WriteHeader(http.StatusBadRequest)
		errorTempalate.Execute(w, "400: Bad Request")
		return
	}

	//removing carriage ret and checking symbols
	var textUpdated string
	for _, symbol := range text {
		if symbol == 13 {
			continue
		}
		if !(symbol >= 32 && symbol <= 126 || symbol == 10) {
			w.WriteHeader(http.StatusBadRequest)
			errorTempalate.Execute(w, "400: Bad Request")
			return
		}
		textUpdated += string(symbol)
	}

	//getting ascii art
	art, err := AsciiArt.GetASCII(textUpdated, font)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorTempalate.Execute(w, "500: Internal Server Error")
		return
	}

	mainTempalate.Execute(w, art)

}

//IndexPage is response to any request excluding ascii art
func IndexPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		errorTempalate.Execute(w, "404: Not found")
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		errorTempalate.Execute(w, "405: Status method not allowed")
		return
	}
	mainTempalate.Execute(w, nil)
}
