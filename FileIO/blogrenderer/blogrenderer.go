package blogrenderer

import (
	"embed"
	"fmt"
	"html/template"
	"io"
)

// temp configuration point read in from file/http eventually
var (
	//go:embed "templates/*"
	postTemplate embed.FS
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func Render(w io.Writer, post Post) error {
	templ, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}
	if err := templ.Execute(w, post); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}
	return nil
}
