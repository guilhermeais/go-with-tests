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
			posts = append(posts, makePostFromFile(fileSystem, file.Name()))
		}

		return posts
	}
	return nil
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
