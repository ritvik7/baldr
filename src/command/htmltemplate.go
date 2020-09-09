package command

import (
	"io/ioutil"
)

func WriteHtmlTemplate(outFile string) error {
	b := []byte(string(templateStr))
	err := ioutil.WriteFile(outFile, b, 0777)
	if err != nil {
		return err
	}
	return nil
}

const templateStr = `
  <html>
  <head>
  <meta charset="utf-8">
  <meta content="IE=edge,chrome=1" http-equiv="X-UA-Compatible">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <title>API Reference</title>

  <link href="/stylesheets/screen.css" rel="stylesheet" type="text/css" media="screen">
  <link href="/stylesheets/print.css" rel="stylesheet" type="text/css" media="print">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js"></script>
  <script src="/javascripts/all.js" type="text/javascript"></script>
  </head>

  <body class="index">
  <a href="#" id="nav-button">
  <span>
  NAV
  <img src="/images/navbar.png">
  </span>
  </a>

  <div class="tocify-wrapper">
  </div>

  <div class="page-wrapper">
	{{.Intro}}
	{{.Auth}}
	{{.Charges}}
	{{.Refunds}}
  </div>
  </body>
  </html>
  `
