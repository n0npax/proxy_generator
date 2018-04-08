package parser

import (
	"io"
	"text/template"
)

// generate nginx config
func GenerateConfig(redirections []NginxRedirection, streamOut io.Writer, templateFile string) {
	tmpl := template.Must(template.ParseFiles(templateFile))

	err := tmpl.Execute(streamOut, redirections)
	if err != nil {
		panic(err)
	}
}

type NginxRedirection struct {
	InternalURL string
	ExternalURL string
}
