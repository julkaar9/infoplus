package server

import (
	"embed"
	"html/template"
)

//go:embed templates/*.html
var embedFS embed.FS

func GetTemplates() (*template.Template, error) {
	allTemplates, err := template.ParseFS(embedFS, "templates/*.html")
	if err != nil {
		return nil, err
	}
	return allTemplates, nil
}
