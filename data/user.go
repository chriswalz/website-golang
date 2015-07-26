package data

import (
	"fmt"
	"math/rand"
	"time"

	"code.google.com/p/go-uuid/uuid"
)

type User struct {
	Uuid, UserName, ParentId, Name, Email, Password, Group string // <-- use as pointer?
	Classrooms                                             []string
	DateCreated                                            time.Time
}

func (user User) String() string {
	return fmt.Sprintln("Name:", user.Uuid, "Email:", user.Email, "Date Created:", user.DateCreated)
}

func CreateRandomUser(d *Dataholder, parentId string) string {
	rand.Seed(5)
	user := User{}
	user.Uuid = uuid.New()
	user.ParentId = parentId
	user.Name = randSeq(rand.Intn(5) + 6)
	user.Email = randSeq(rand.Intn(7) + 8)
	user.DateCreated = time.Now()
	// group can either be "student" or "teacher"
	d.AddUser(user)
	return user.Uuid
}
func (user User) AddRandomClassroomList(d *Dataholder) {

	for i := 0; i < 2; i++ {
		user.Classrooms = append(user.Classrooms, CreateRandomClassroom(d, user.Uuid))
	}
}
