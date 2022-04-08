package entity

import "time"

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}
