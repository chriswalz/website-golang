// main.go
package main

/*
Notes:
Create___ may need to return two values: UUID and struct
Add all createRandomList methods
Create tests, initialize a database
Implement post, get, put actions
Connect RestApi
*/

import (
	"fmt"
	"time"

	"github.com/website-golang/data"
)

func main() {
	defer data.TimeTrack(time.Now(), "Main")
	var d = data.CreateRandomDatabase()
	d.PrintDataHolder()
	//fmt.Println("Users:", d.Users)
	//fmt.Println("Classrooms:", d.Classrooms)
	fmt.Println("Questions length:", len(d.Questions))

}
