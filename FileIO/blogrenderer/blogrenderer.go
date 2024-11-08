package blogrenderer

import (
	"embed"
	"errors"
	"fmt"
	"html/template"
	"io"
	"strings"
)

// temp configuration point read in from file/http eventually
var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

// improve benchmarks
type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, errors.New("error parsing templates from filesys")
	}
	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

func (r *PostRenderer) Render(w io.Writer, post Post) error {
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", post); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}
	return nil
}
