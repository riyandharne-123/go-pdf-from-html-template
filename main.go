package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"

	u "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type User struct {
	Name string
}

func main() {
	u.SetPath("C:\\Program Files\\wkhtmltopdf\\bin\\wkhtmltopdf.exe")

	// Create new PDF generator
	pdfg, err := u.NewPDFGenerator()
	if err != nil {
		panic(err)
	}

	//Set global options
	user := User{
		Name: "Riyan Dharne",
	}

	//set template path
	html, err := ioutil.ReadFile("templates/test.html")
	if err != nil {
		panic(err)
	}

	//parse html template
	htmlString := string(html)
	htmlTemplate, err := template.New("test").Parse(htmlString)
	if err != nil {
		panic(err)
	}

	//set variables and convert to string
	var htmlTemplateString bytes.Buffer
	err = htmlTemplate.Execute(&htmlTemplateString, user)
	if err != nil {
		panic(err)
	}

	//add page
	pdfg.AddPage(u.NewPageReader(strings.NewReader(htmlTemplateString.String())))

	//Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		panic(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./storage/test.pdf")
	if err != nil {
		panic(err)
	}

	fmt.Println("Done")
}
