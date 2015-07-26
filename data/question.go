package data

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"math/rand"
	"time"
)

type Question struct {
	Uuid, Name, ParentId, CreatorId string // <-- use as pointer?
	upvotes                         int
	DateCreated                     time.Time
}

func (q Question) PrintQuestion() { // Question name is the actual question
	fmt.Println("Name:", q.Name, "Date Created:", q.DateCreated)
}

func CreateRandomQuestion(d *Dataholder, parentId string) string { // parent should be a classsroom
	rand.Seed(6)
	q := Question{}
	q.ParentId = parentId
	q.Uuid = uuid.New()
	q.Name = randSeq(rand.Intn(5)+3) + "?"
	q.DateCreated = time.Now()
	d.AddQuestion(q)
	return q.Uuid
}
