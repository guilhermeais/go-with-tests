package blogposts_test

import (
	"errors"
	"io/fs"
	blogposts "reading-files"
	"reflect"
	"testing"
	"testing/fstest"
	"time"
)

func TestPostsFromFS(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		fileSystem := fstest.MapFS{
			"hello-world.md": {Data: []byte("Title: Hello, TDD world!\nDescription: First post on our wonderful blog\nTags: tdd, go\nDate: 2022-08-26\n---\nHello world!")},
		}

		posts, err := blogposts.PostsFromFS(fileSystem)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fileSystem) {
			t.Errorf("expected %d posts, got %d posts", len(fileSystem), len(posts))
		}

		expectedFirstPost := blogposts.Post{
			Title:       "Hello, TDD world!",
			Description: "First post on our wonderful blog",
			Tags:        []string{"tdd", "go"},
			Body:        "Hello world!",
			Date:        time.Date(2022, 8, 26, 0, 0, 0, 0, time.UTC),
		}

		assertPost(t, posts[0], expectedFirstPost)
	})

	t.Run("sort the posts by date", func(t *testing.T) {
		fileSystem := fstest.MapFS{
			"hello-world.md": {Data: []byte("Title: Hello, TDD world!\nDescription: First post on our wonderful blog\nTags: tdd, go\nDate: 2022-08-26\n---\nHello world!")},
			"other post.md":  {Data: []byte("Title: God is good!\nDescription: a post about how god is good\nTags: god, good\nDate: 2022-08-25\n---\nGod is good!")},
		}

		posts, err := blogposts.PostsFromFS(fileSystem)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fileSystem) {
			t.Errorf("expected %d posts, got %d posts", len(fileSystem), len(posts))
		}

		expectedFirstPost := blogposts.Post{
			Title:       "God is good!",
			Description: "a post about how god is good",
			Tags:        []string{"god", "good"},
			Body:        "God is good!",
			Date:        time.Date(2022, 8, 25, 0, 0, 0, 0, time.UTC),
		}

		expectedSecondPost := blogposts.Post{
			Title:       "Hello, TDD world!",
			Description: "First post on our wonderful blog",
			Tags:        []string{"tdd", "go"},
			Body:        "Hello world!",
			Date:        time.Date(2022, 8, 26, 0, 0, 0, 0, time.UTC),
		}

		assertPost(t, posts[0], expectedFirstPost)
		assertPost(t, posts[1], expectedSecondPost)
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

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

type FailingFS struct {
}

func (f FailingFS) Open(_ string) (fs.File, error) {
	return nil, errors.New("i've failed")
}
