package blogposts_test

import (
	blogposts "reading-files"
	"testing"
	"testing/fstest"
)

func TestBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte("hello, world")},
		"hello-twitch.md": {Data: []byte("hello, twitch")},
	}

	posts := blogposts.PostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("expected %d posts, got %d posts", len(fs), len(posts))
	}
}
