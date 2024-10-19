package blogposts

import (
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title string
}

func PostsFromFS(fileSystem fs.FS) []Post {
	dir, err := fs.ReadDir(fileSystem, ".")
	posts := []Post{}
	if err == nil {
		for _, file := range dir {
			posts = append(posts, makePostFromFile(fileSystem, file))
		}

		return posts
	}
	return nil
}

func makePostFromFile(fileSystem fs.FS, file fs.DirEntry) Post {
	blogFile, _ := fileSystem.Open(file.Name())
	fileContents, _ := io.ReadAll(blogFile)

	title := strings.TrimPrefix(string(fileContents), "Title: ")
	return Post{Title: title}
}
