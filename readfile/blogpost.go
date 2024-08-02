package blogpost

import "io/fs"

func NewPostsFromFS(fileSys fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSys, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, file := range dir {
		post, err := getPost(fileSys, file.Name())
		if err != nil {
			return nil, err //todo: clarification
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSys fs.FS, fileName string) (Post, error) {
	postFile, err := fileSys.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}
