package main

import (
	"log"
	"os"

	bp "example.com/blogposts"
)

func main() {
	posts, err := bp.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatalf("NewPostsFromFS err: %v", err)
	}
	for _, post := range posts {
		log.Printf("Post: %+v\n", post)
	}
}
