package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/foo", fooHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9999" // Default port if not specified
	}
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the home page!!!")
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	ExampleNewPDFGenerator()
	fmt.Fprint(w, "<html><body><h1>FOOOOOOOO</h1></body></html>")
}

// A funciton to handle a POST request with a JSON payload as the body
func postHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body

	// Decode the JSON payload
	// Do something with the data
	// Return a response
}

func ExampleNewPDFGenerator() {

	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	pdfg.Grayscale.Set(true)

	// Create a new input page from an URL
	page := wkhtmltopdf.NewPage("https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf")

	// Set options for this page
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(0.95)

	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./simplesample.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}
