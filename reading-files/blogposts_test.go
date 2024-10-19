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

		assertPost(t, posts[0], expectedFirstPost)
	})

	t.Run("failing filesystem", func(t *testing.T) {
		_, err := blogposts.PostsFromFS(FailingFS{})

		if err == nil {
			t.Error("expected an error, dignt get one")
		}
	})
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()

	if got != want {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

type FailingFS struct {
}

func (f FailingFS) Open(_ string) (fs.File, error) {
	return nil, errors.New("i've failed")
}
