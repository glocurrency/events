package auth

import "time"

type UserRegistered struct {
	ID   string
	Date time.Time
}
