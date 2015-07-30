package data

import (
	"fmt"
	"math/rand"
	"time"

	"code.google.com/p/go-uuid/uuid"
)

var Db *Dataholder

type Dataholder struct {
	Uuid, Name  string // <-- use as pointer?
	Users       []*User
	Classrooms  []*Classroom
	Questions   []*Question
	Sessions    map[string]*Session
	DateCreated time.Time
}

func (d *Dataholder) PrintDataHolder() {
	fmt.Println("Name:", d.Name, "Date Created:", d.DateCreated)
}
func (d *Dataholder) AddUser(u *User) {
	d.Users = append(d.Users, u)
}
func (d *Dataholder) AddClassroom(c *Classroom) {
	d.Classrooms = append(d.Classrooms, c)
}
func (d *Dataholder) AddQuestion(q *Question) {
	d.Questions = append(d.Questions, q)
}
func AddSession(user *User) {
	s := CreateSession(user)
	Db.Sessions[s.Token] = s
	//return &d // edit Create Database
}

// create save function and load function
func createRandomDataHolder() *Dataholder {
	rand.Seed(7)
	d := Dataholder{}
	d.Sessions = make(map[string]*Session)
	d.Uuid = uuid.New()
	d.Name = randSeq(rand.Intn(5) + 6)
	d.DateCreated = time.Now()
	return &d // edit Create Database
}
func (d *Dataholder) AddRandomUserList() {
	defer TimeTrack(time.Now(), "Users")
	userc := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			CreateRandomUser(d)
			userc <- i
		}()

	}

	for i := 0; i < 10; i++ {
		<-userc
		//fmt.Println(g)
	}
}
func (d *Dataholder) AddAllClassrooms() {
	defer TimeTrack(time.Now(), "Classrooms")
	for i := 0; i < len(d.Users); i++ {
		d.Users[i].AddRandomClassroomList()
	}
}
func (d *Dataholder) AddAllQuestions() {
	defer TimeTrack(time.Now(), "Questions")
	for i := 0; i < len(d.Classrooms); i++ {
		d.Classrooms[i].AddRandomQuestionList()
	}
}

func CreateRandomDatabase() *Dataholder {
	Db = createRandomDataHolder()
	Db.AddRandomUserList()
	Db.AddAllClassrooms()
	Db.AddAllQuestions()
	return Db
}

/*
	var d = data.CreateRandomDatabase()
	d.PrintDataHolder()
	//fmt.Println("Users:", d.Users)
	//fmt.Println("Classrooms:", d.Classrooms)
	fmt.Println("Questions length:", len(d.Questions))
*/
