package data

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"math/rand"
	"time"
)

type Classroom struct {
	Uuid, Name, ParentId, CreatorId string // <-- use as pointer?
	Questions                       []string
	UserIds                         []string
	DateCreated                     time.Time
}

func (c Classroom) PrintClassRoom() {
	fmt.Println("Name:", c.Name, "Date Created:", c.DateCreated)
}

func CreateRandomClassroom(parentId string) Classroom {
	rand.Seed(6)
	c := Classroom{}
	c.ParentId = parentId
	c.UserIds = append(c.UserIds, parentId) // create will be first user on list
	c.Uuid = uuid.New()
	c.Name = randSeq(rand.Intn(5) + 6)
	c.DateCreated = time.Now()
	return c
}
