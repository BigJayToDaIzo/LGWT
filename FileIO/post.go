package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	titleSep = "Title: "
	descSep  = "Description: "
	tagsSep  = "Tags: "
	tagBod   = "Body: "
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func newPost(postBody io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postBody)
	readMetaLine := func(key string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), key)
	}
	title := readMetaLine(titleSep)
	desc := readMetaLine(descSep)
	tags := strings.Split(readMetaLine(tagsSep), ", ")
	scanner.Scan() // ignore the --- we end metadata with
	// parse body with buffer
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	body := strings.TrimSuffix(buf.String(), "\n")

	return Post{
		Title:       title,
		Description: desc,
		Tags:        tags,
		Body:        body,
	}, nil
}
