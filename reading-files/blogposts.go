package blogposts

import (
	"io/fs"
)

type Post struct{}

func PostsFromFS(fileSysten fs.FS) []Post {
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
