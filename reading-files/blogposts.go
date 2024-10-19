package blogposts

import (
	"io"
	"io/fs"
	"strings"
)

func PostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	posts := []Post{}
	for _, file := range dir {
		posts = append(posts, makePostFromFile(fileSystem, file.Name()))
	}

	return posts, nil
}

func makePostFromFile(fileSystem fs.FS, filename string) Post {
	blogFile, _ := fileSystem.Open(filename)
	return newPost(blogFile)
}

func newPost(blogReader io.Reader) Post {
	fileContents, _ := io.ReadAll(blogReader)

	title := strings.TrimPrefix(string(fileContents), "Title: ")
	return Post{Title: title}
}
