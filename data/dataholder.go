package data

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"math/rand"
	"time"
)

type Dataholder struct {
	Uuid, Name  string // <-- use as pointer?
	Users       []User
	Classrooms  []Classroom
	Questions   []Question
	DateCreated time.Time
}

func (d Dataholder) PrintDataHolder() {
	fmt.Println("Name:", d.Name, "Date Created:", d.DateCreated)
}

func CreateRandomDataHolder() Dataholder {
	rand.Seed(7)
	d := Dataholder{}
	d.Uuid = uuid.New()
	d.Name = randSeq(rand.Intn(5) + 6)
	d.DateCreated = time.Now()
	return d
}
func (d Dataholder) AddRandomUserList() {
	for i := 0; i < 1000; i++ {
		d.Users = append(d.Users, CreateRandomUser(d.Uuid))
	}
}
