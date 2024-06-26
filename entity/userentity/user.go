package userentity

import "time"

type User struct {
	ID        string
	FirstName string
	LastName  string
	Password  string
	CreatedAt time.Time
}
