package command

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type HeaderJson struct {
	Swagger string
	Info    struct {
		Version   string
		Title     string
		Languages []string
	}
}

type Json struct {
	Info struct {
		Description    string
		Authentication struct {
			Summary  string
			Snippets struct {
				C      string
				Cpp    string
				Erlang string
				Go     string
			}
			Note string
		}
		Termsofservice string
		Contact        struct {
			Name string
			URL  string
		}
		License struct {
			Name string
			URL  string
		}
	}
	Paths struct {
		Charges struct {
			Get struct {
				Summary string `json:"summary"`
			} `json:"get"`
		} `json:"/charges"`
		Refunds struct {
			Get struct {
				Summary string `json:"summary"`
			} `json:"get"`
		} `json:"/refunds"`
	} `json:"paths"`
}

const (
	DIR  = iota
	FILE = iota
)

type AssetInfo struct {
	Path string
	Type int
}

func getCodeSnippet(url string) string {
	s := ""
	s = s + "```\n"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	s = s + string(body)
	s = s + "\n```"
	return s
}
func createSiteStructure(rootPath string) {
	p := [19]AssetInfo{
		{"", DIR},
		{"/source", DIR},
		{"/source/assets", DIR},
		{"/source/assets/fonts", DIR},
		{"/source/assets/fonts", DIR},
		{"/source/assets/javascripts", DIR},
		{"/source/assets/images", DIR},
		{"/source/assets/stylesheets", DIR},
		{"/source/markdowns", DIR},
		{"/source/layouts", DIR},
		{"/source/layouts/index.html.template", FILE},
		{"/source/markdowns/intro.md", FILE},
		{"/source/markdowns/authentication.md", FILE},
		{"/source/markdowns/header.json", FILE},
		{"/source/markdowns/charges.md", FILE},
		{"/source/markdowns/refunds.md", FILE},
		{"/dist", DIR},
		{"/README.md", FILE},
		{"/settings.json", FILE},
	}

	for i := range p {
		p[i].Path = string(rootPath + p[i].Path)
		if p[i].Type == DIR {
			os.MkdirAll(p[i].Path, 0777)
		} else {
			os.Create(p[i].Path)
		}
	}
}

func initc(root string, path string) {
	if path == "" {
		fmt.Printf("Path not specified!\n")
		os.Exit(1)
	}
	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	dec := json.NewDecoder(strings.NewReader(string(file)))
	var d Json
	var c HeaderJson
	for {
		err := dec.Decode(&d)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Error!")
		}
	}
	dec = json.NewDecoder(strings.NewReader(string(file)))
	for {
		err := dec.Decode(&c)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Error!")
		}
	}
	createSiteStructure(root)

	WriteCSSAsset(root+"/source/assets/stylesheets/screen.css", "screen")
	WriteCSSAsset(root+"/source/assets/stylesheets/print.css", "print")

	WriteJSAssetAll(root + "/source/assets/javascript/all.js")

	h, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile(root+"/source/markdowns/header.json", h, 0777)

	intro := IntroMd{d.Info.Description}
	WriteMdSection(root+"/source/markdowns/intro.md", intro, "introduction")
	auth := AuthenticationMd{
		d.Info.Authentication.Summary,
		getCodeSnippet(d.Info.Authentication.Snippets.C),
		getCodeSnippet(d.Info.Authentication.Snippets.Cpp),
		getCodeSnippet(d.Info.Authentication.Snippets.Erlang),
		getCodeSnippet(d.Info.Authentication.Snippets.Go),
		d.Info.Authentication.Note,
	}
	WriteMdSection(root+"/source/markdowns/authentication.md", auth, "authentication")
	entity := EntityMd{"Charges", d.Paths.Charges.Get.Summary}
	WriteMdSection(root+"/source/markdowns/charges.md", entity, "entity")
	entity = EntityMd{"Refunds", d.Paths.Refunds.Get.Summary}
	WriteMdSection(root+"/source/markdowns/refunds.md", entity, "entity")
	WriteHtmlTemplate(root + "/source/layouts/index.html.template")
	fmt.Printf("Created!\n")
}

var Init = cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "Create a new api doc site from the command line.",
	Action: func(c *cli.Context) {
		println("Creating new site.....", c.Args()[0])
		initc(c.Args()[0], c.Args()[1])
	},
}
