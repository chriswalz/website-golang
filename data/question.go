package data

import (
	"fmt"
	"math/rand"
	"time"

	"code.google.com/p/go-uuid/uuid"
)

type Question struct {
	Uuid, Name, CreatorId string // <-- use as pointer?
	upvotes               int
	Parent                *Classroom
	DateCreated           time.Time
}

func (q *Question) PrintQuestion() { // Question name is the actual question
	fmt.Println("Name:", q.Name, "Date Created:", q.DateCreated)
}

func CreateRandomQuestion(parent *Classroom) *Question { // parent should be a classsroom
	rand.Seed(6)
	q := Question{}
	q.Parent = parent
	q.Uuid = uuid.New()
	q.Name = randSeq(rand.Intn(5)+3) + "?"
	q.DateCreated = time.Now()
	Db.AddQuestion(&q)
	return &q
}
