package server

import (
	"time"
)

type Post struct {
	ID         int
	Source     string
	Content    string
	CreatedAt  time.Time
	ModifiedAt time.Time
}
