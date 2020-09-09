package command

import (
	"bytes"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/russross/blackfriday"
	"html"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type build struct {
	Intro   string
	Auth    string
	Charges string
	Refunds string
}

var d build

func BuildTemplate(outFile string) {
	templ, err := template.New("index.html.template").ParseFiles("source/layouts/index.html.template")
	if err != nil {
		fmt.Println(err)
	}

	path := "source/markdowns/"

	a, e := ioutil.ReadFile(path + "intro.md")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	h := blackfriday.MarkdownCommon(a)
	d.Intro = string(h)

	a, e = ioutil.ReadFile(path + "authentication.md")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	h = blackfriday.MarkdownCommon(a)
	d.Auth = string(h)

	a, e = ioutil.ReadFile(path + "charges.md")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	h = blackfriday.MarkdownCommon(a)
	d.Charges = string(h)

	a, e = ioutil.ReadFile(path + "refunds.md")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	h = blackfriday.MarkdownCommon(a)
	d.Refunds = string(h)

	b := new(bytes.Buffer)
	err = templ.Execute(b, d)
	if err != nil {
		fmt.Println(err)
	}
	s := b.String()
	w := html.UnescapeString(s)
	err = ioutil.WriteFile(outFile, []byte(w), 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func buildc() {
	src := "source/assets"
	dest := "dist"
	err := CopyDir(src, dest)
	if err != nil {
		fmt.Println("Error")
	} else {
		log.Print("Files copied.")
	}
	os.Create("dist/index.html")
	BuildTemplate("dist/index.html")
}

var Build = cli.Command{
	Name:    "build",
	Aliases: []string{"b"},
	Usage:   "Builds the static site ready for production deployment.",
	Action: func(c *cli.Context) {
		println("Building Documentation Site...")
		buildc()
		println("Done!")
	},
}
