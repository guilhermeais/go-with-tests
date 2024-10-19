package blogposts_test

import (
	blogposts "reading-files"
	"testing"
	"testing/fstest"
)

func TestBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"1-hello-world.md": {Data: []byte("Title: Hello, TDD world!")},
		"hello-twitch.md":  {Data: []byte("Title: Hello, twitchy world!")},
	}

	posts := blogposts.PostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("expected %d posts, got %d posts", len(fs), len(posts))
	}

	expectedFirstPost := blogposts.Post{Title: "Hello, TDD world!"}
	if posts[0] != expectedFirstPost {
		t.Errorf("got %#v, want %#v", posts[0], expectedFirstPost)
	}
}
