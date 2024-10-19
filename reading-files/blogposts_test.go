package blogposts_test

import (
	"errors"
	"io/fs"
	blogposts "reading-files"
	"testing"
	"testing/fstest"
)

func TestPostsFromFS(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		fs := fstest.MapFS{
			"1-hello-world.md": {Data: []byte("Title: Hello, TDD world!\nDescription: First post on our wonderful blog")},
			"hello-twitch.md":  {Data: []byte("Title: Hello, twitchy world!")},
		}

		posts, err := blogposts.PostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("expected %d posts, got %d posts", len(fs), len(posts))
		}

		expectedFirstPost := blogposts.Post{
			Title:       "Hello, TDD world!",
			Description: "First post on our wonderful blog",
		}

		if posts[0] != expectedFirstPost {
			t.Errorf("got %#v, want %#v", posts[0], expectedFirstPost)
		}
	})

	t.Run("failing filesystem", func(t *testing.T) {
		_, err := blogposts.PostsFromFS(FailingFS{})

		if err == nil {
			t.Error("expected an error, dignt get one")
		}
	})
}

type FailingFS struct {
}

func (f FailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("i've failed")
}
