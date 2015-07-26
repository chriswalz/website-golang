package data

import (
	"fmt"
	"math/rand"
	"time"

	"code.google.com/p/go-uuid/uuid"
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

func CreateRandomClassroom(d *Dataholder, parentId string) string {
	rand.Seed(6)
	c := Classroom{}
	c.ParentId = parentId
	c.UserIds = append(c.UserIds, parentId) // create will be first user on list
	c.Uuid = uuid.New()
	c.Name = randSeq(rand.Intn(5) + 6)
	c.DateCreated = time.Now()
	d.AddClassroom(c)
	return c.Uuid
}
func (c Classroom) AddRandomQuestionList(d *Dataholder) {
	for i := 0; i < 2; i++ {
		c.Questions = append(c.Questions, CreateRandomQuestion(d, c.Uuid))
	}
}
