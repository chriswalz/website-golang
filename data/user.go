package data

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	Uuid, UserName, ParentId, Name, Email, Password, Group string // <-- use as pointer?
	Classrooms                                             []string
	DateCreated                                            time.Time
}

func (user User) PrintUser() {
	fmt.Println("Name:", user.Name, "Email:", user.Email, "Date Created:", user.DateCreated)
}

func CreateRandomUser(parentId string) User {
	rand.Seed(5)
	user := User{}
	user.Uuid = uuid.New()
	user.ParentId = parentId
	user.Name = randSeq(rand.Intn(5) + 6)
	user.Email = randSeq(rand.Intn(7) + 8)
	user.DateCreated = time.Now()
	// group can either be "student" or "teacher"
	return user
}
func (user User) AddRandomClassroomList() {
	for i := 0; i < 2; i++ {
		user.Classrooms = append(user.Classrooms, CreateRandomClassroom(user.Uuid))
	}
}
