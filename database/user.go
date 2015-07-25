package database

import (
	"code.google.com/p/go-uuid/uuid"
	"time"
)

type User struct {
	uuid        uuid.UUID
	name        string // <-- use as pointer?
	email       string
	classrooms  []uuid.UUID
	dateCreated time.Time
}
