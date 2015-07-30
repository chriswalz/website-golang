package data

import (
	"fmt"
	"math/rand"
	"time"

	"code.google.com/p/go-uuid/uuid"
)

type Classroom struct {
	Uuid, Name, CreatorId string // <-- use as pointer?
	Questions             []*Question
	Users                 []*User
	Parent                *User
	DateCreated           time.Time
}

func (c *Classroom) PrintClassRoom() {
	fmt.Println("Name:", c.Name, "Date Created:", c.DateCreated)
}

func CreateRandomClassroom(parent *User) *Classroom {
	rand.Seed(6)
	c := Classroom{}
	c.Parent = parent
	c.Users = append(c.Users, parent) // create will be first user on list
	c.Uuid = uuid.New()
	c.Name = randSeq(rand.Intn(5) + 6)
	c.DateCreated = time.Now()
	Db.AddClassroom(&c)
	return &c
}
func (c *Classroom) AddRandomQuestionList() {
	for i := 0; i < 2; i++ {
		c.Questions = append(c.Questions, CreateRandomQuestion(c))
	}
}
