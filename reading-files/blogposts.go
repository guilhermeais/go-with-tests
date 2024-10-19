package blogposts

import (
	"bufio"
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
	scanner := bufio.NewScanner(blogReader)

	scanner.Scan()
	title := strings.TrimPrefix(scanner.Text(), "Title: ")

	scanner.Scan()
	description := strings.TrimPrefix(scanner.Text(), "Description: ")

	return Post{
		Title:       title,
		Description: description,
	}
}
