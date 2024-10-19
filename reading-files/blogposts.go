package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"sort"
	"strings"
	"time"
)

const (
	TITLE_PREFIX       = "Title: "
	DESCRIPTION_PREFIX = "Description: "
	TAG_PREFIX         = "Tags: "
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

	sort.Slice(posts, func(i, j int) bool {
		a := posts[i]
		b := posts[j]
		return a.Date.Before(b.Date)
	})

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
	tags := strings.Split(readLine(TAG_PREFIX), ", ")
	dateStr := readLine("Date: ")
	date, _ := time.Parse("2006-01-02", dateStr)

	body := readBody(scanner)
	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        body,
		Date:        date,
	}
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()
	var buf bytes.Buffer
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	return strings.TrimSuffix(buf.String(), "\n")
}
