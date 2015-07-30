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
create pointers to questions user owns, do equivalent for classrooms
database should be an array of pointers
need to remove invalid sessions after 30 minutes
*/

import (
	"time"

	"github.com/website-golang/data"
	"github.com/website-golang/handlers"
)

func main() {
	defer data.TimeTrack(time.Now(), "Main")
	data.CreateRandomDatabase()
	handlers.Prepare()
}
func rest() {
	// Middlewares: gzip, authbasic
	// others: logging,
	//rest.App
}
