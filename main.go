// main.go
package main

/*
Notes:
Finish reading "Writing Web apps in Go"
Create___ may need to return two values: UUID and struct
Add all createRandomList methods
Create tests, initialize a database
Implement post, get, put actions
Connect RestApi
Add persistent storage
*/

import (
	"time"

	"github.com/website-golang/data"
	"github.com/website-golang/handlers"
)

func main() {
	defer data.TimeTrack(time.Now(), "Main")
	handlers.Prepare()
}
func rest() {
	// Middlewares: gzip, authbasic
	// others: logging,
	//rest.App
}

//defer data.TimeTrack(time.Now(), "Main")
/*
	var d = data.CreateRandomDatabase()
	d.PrintDataHolder()
	//fmt.Println("Users:", d.Users)
	//fmt.Println("Classrooms:", d.Classrooms)
	fmt.Println("Questions length:", len(d.Questions))
*/
