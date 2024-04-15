package auth

import "time"

type UserRegistered struct {
	ID        string
	Provider  string
	Timestamp time.Time
}
