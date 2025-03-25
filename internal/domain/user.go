package domain

import "time"

type UserID int64
type User struct {
	ID        UserID
	ClientID  ClientID
	Name      string
	CreatedAt time.Time
}
