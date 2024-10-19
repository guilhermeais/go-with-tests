package blogposts

import "time"

type Post struct {
	Title, Description, Body string
	Tags                     []string
	Date                     time.Time
}
