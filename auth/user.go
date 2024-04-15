package auth

import "time"

const UserRegisteredEvent = "user-registered"

type UserRegistered struct {
	ID        string
	Provider  string
	Timestamp time.Time
}
