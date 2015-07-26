package data

import (
	"fmt"
	"math/rand"
	"time"

	"code.google.com/p/go-uuid/uuid"
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
func (d *Dataholder) AddUser(u User) {
	d.Users = append(d.Users, u)
}
func (d *Dataholder) AddClassroom(c Classroom) {
	d.Classrooms = append(d.Classrooms, c)
}
func (d *Dataholder) AddQuestion(q Question) {
	d.Questions = append(d.Questions, q)
}
func createRandomDataHolder() *Dataholder {
	rand.Seed(7)
	d := Dataholder{}
	d.Uuid = uuid.New()
	d.Name = randSeq(rand.Intn(5) + 6)
	d.DateCreated = time.Now()
	return &d
}
func (d *Dataholder) AddRandomUserList() {
	defer TimeTrack(time.Now(), "Users")
	userc := make(chan int)

	for i := 0; i < 30000; i++ {
		go func() {
			CreateRandomUser(d, d.Uuid)
			userc <- i
		}()

	}

	for i := 0; i < 30000; i++ {
		<-userc
		//fmt.Println(g)
	}
}
func (d *Dataholder) AddAllClassrooms() {
	defer TimeTrack(time.Now(), "Classrooms")
	for i := 0; i < len(d.Users); i++ {
		d.Users[i].AddRandomClassroomList(d)
	}
}
func (d *Dataholder) AddAllQuestions() {
	defer TimeTrack(time.Now(), "Questions")
	for i := 0; i < len(d.Classrooms); i++ {
		d.Classrooms[i].AddRandomQuestionList(d)
	}
}

func CreateRandomDatabase() Dataholder {
	d := createRandomDataHolder()
	d.AddRandomUserList()
	d.AddAllClassrooms()
	d.AddAllQuestions()
	return *d

}
