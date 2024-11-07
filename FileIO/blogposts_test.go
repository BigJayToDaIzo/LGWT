package blogposts_test

import (
	"errors"
	"fmt"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	b "example.com/blogposts"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("woopsie doosie! fileSys failure error thrown")
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)
	t.Run("happy path fs read", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := b.NewPostsFromFS(fs)
		if err != nil {
			t.Fatal(fmt.Printf("ReadPostFromFS err: %v", err))
		}
		assertPost(t, posts[0], b.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World`,
		})
	})
	t.Run("sad path fs read", func(t *testing.T) {
		_, err := b.NewPostsFromFS(StubFailingFS{})
		if err == nil {
			t.Fatal("expected a fs read error but got none")
		}
	})
}

func assertPost(t *testing.T, got b.Post, want b.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
}
