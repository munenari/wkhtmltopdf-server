package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus * 2)

	var bind int
	flag.IntVar(&bind, "p", 10000, "bind web server port. default: 10000")
	flag.Parse()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.File("/", "html/index.html")
	e.POST("/gen", generateAction)
	e.HideBanner = true
	log.Printf("server started with port: %d\n", bind)
	log.Fatalln(e.Start(fmt.Sprintf(":%d", bind)))
}

func generateAction(c echo.Context) error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request().Body)
	pdfbuf, err := generate(buf)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.Stream(http.StatusOK, "application/pdf", pdfbuf)
}

func generate(buf *bytes.Buffer) (*bytes.Buffer, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	pdfg.AddPage(wkhtmltopdf.NewPageReader(buf))
	err = pdfg.Create()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return pdfg.Buffer(), nil
}
