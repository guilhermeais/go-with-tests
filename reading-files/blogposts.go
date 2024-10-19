package blogposts

import (
	"bufio"
	"io"
	"io/fs"
	"strings"
)

const (
	TITLE_PREFIX       = "Title: "
	DESCRIPTION_PREFIX = "Description: "
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

	readLine := func(prefix string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), prefix)
	}

	title := readLine(TITLE_PREFIX)
	description := readLine(DESCRIPTION_PREFIX)

	return Post{
		Title:       title,
		Description: description,
	}
}
