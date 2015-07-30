package data

import (
	"time"

	"code.google.com/p/go-uuid/uuid"
)

type Session struct {
	DateCreated time.Time
	Token       string
	Parent      *User
	Expiration  time.Time
}

func CreateSession(parent *User) *Session {
	s := Session{}
	s.DateCreated = time.Now()
	s.Token = uuid.New()
	s.Parent = parent
	s.Expiration = s.DateCreated.Add(time.Minute * 30)
	return &s
}
