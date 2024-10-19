package blogposts

import (
	"io/fs"
	"testing/fstest"
)

type Post struct{}

func PostsFromFS(fileSysten fstest.MapFS) []Post {
	dir, err := fs.ReadDir(fileSysten, ".")
	posts := []Post{}
	if err == nil {
		for range dir {
			posts = append(posts, Post{})
		}

		return posts
	}
	return nil
}
