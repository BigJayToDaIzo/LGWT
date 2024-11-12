package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	blr "example.com/blogposts/blogrenderer"
	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	var (
		aPost = blr.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	postRenderer, err := blr.NewPostRenderer()
	if err != nil {
		t.Fatal("unexpected error creating post renderer", err)
	}
	t.Run("it converts a post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal("unexpected error converting post to HTML", err)
		}
		approvals.VerifyString(t, buf.String())
	})
	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blr.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}
		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal("unexpected error rendering index", err)
		}
		got := buf.String()
		want := `<ol><li><a href="/post/hello-world">Hello World</a></li><li><a href="/post/hello-world-2">Hello World 2</a></li></ol>
`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func BenchmarkRenderer(b *testing.B) {
	var (
		aPost = blr.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "Scription",
			Tags:        []string{"go", "tdd"},
		}
	)
	postRenderer, err := blr.NewPostRenderer()
	if err != nil {
		b.Fatal("unexpected error creating post renderer", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
