package blogposts

import (
	"errors"
	"io/fs"
)

// loosen coupleing from fs.FS to io.Reader
func NewPostsFromFS(fileSys fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSys, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, file := range dir {
		post, postErr := getPost(fileSys, file.Name())
		if postErr != nil {
			return nil, errors.New("error reading post")
		}
		posts = append(posts, post)
	}
	return posts, nil

}

func getPost(fileSys fs.FS, fileName string) (Post, error) {
	postFile, err := fileSys.Open(fileName)
	if err != nil {
		return Post{}, err // TODO Do gud'r'n this
	}
	defer postFile.Close()

	return newPost(postFile)
}
