package data

import (
	"fmt"
	"math/rand"
	"time"

	"code.google.com/p/go-uuid/uuid"
)

type User struct {
	Uuid, UserName, Name, Email, Password, Group string // <-- use as pointer?
	Classrooms                                   []*Classroom
	Parent                                       *Dataholder
	DateCreated                                  time.Time
}

func IsUnavailable(userName string) bool {
	for _, user := range Db.Users {
		if user.UserName == userName {
			return true
		}
	}
	return false
}
func (user *User) String() string {
	return fmt.Sprintln("Name:", user.Uuid, "Email:", user.Email, "Date Created:", user.DateCreated)
}
func CreateUser(parent *Dataholder, username, password, name, email string) *User {
	user := User{}
	user.Uuid = uuid.New()
	user.Parent = parent
	user.UserName = username
	user.Password = password
	user.Name = name
	user.Email = email
	user.DateCreated = time.Now()
	// group can either be "student" or "teacher"
	Db.AddUser(&user)
	return &user
}
func CreateRandomUser(parent *Dataholder) *User {
	rand.Seed(5)
	user := User{}
	user.Uuid = uuid.New()
	user.Parent = parent
	user.Name = randSeq(rand.Intn(5) + 6)
	user.Email = randSeq(rand.Intn(7) + 8)
	user.DateCreated = time.Now()
	// group can either be "student" or "teacher"
	Db.AddUser(&user)
	return &user
}
func (user *User) AddRandomClassroomList() {

	for i := 0; i < 2; i++ {
		user.Classrooms = append(user.Classrooms, CreateRandomClassroom(user))
	}
}
