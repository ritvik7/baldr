package command

import (
	"bytes"
	"io/ioutil"
	"text/template"
)

var sectionMap = map[string]string{
	"introduction":   introductionMD,
	"authentication": authenticationMD,
	"entity":         entityMD,
}

// Section data structs ...
type IntroMd struct {
	Content string
}

type AuthenticationMd struct {
	Summary string
	C       string
	Cpp     string
	Erlang  string
	Go      string
	Note    string
}

type EntityMd struct {
	Entity      string
	Description string
}

func WriteMdSection(outFile string, sectionData interface{}, sectionName string) error {
	tmpl, err := template.New(sectionName).Parse(sectionMap[sectionName])
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, sectionData)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outFile, []byte(buf.String()), 0777)

	if err != nil {
		return err
	}

	return nil
}

// Markdown templates for sections ....

const introductionMD = `
# Introduction

{{.Content}}
`

const authenticationMD = `
# Authentication

{{.Summary}}

- C
{{.C}}
- Cpp
{{.Cpp}}
- Erlang
{{.Erlang}}
- Go
{{.Go}}

<aside class="note">
{{.Note}}
</aside>
`

const entityMD = `
# {{.Entity}}

## {{.Description}}
`
