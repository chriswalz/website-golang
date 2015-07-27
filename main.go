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
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/website-golang/data"
	"github.com/website-golang/file"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	f, _ := file.Load("html/weeb.html")
	fmt.Fprintf(w, string(f.Body), r.URL.Path[1:])
}
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}
func main() {
	defer data.TimeTrack(time.Now(), "Main")
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}

//defer data.TimeTrack(time.Now(), "Main")
/*
	var d = data.CreateRandomDatabase()
	d.PrintDataHolder()
	//fmt.Println("Users:", d.Users)
	//fmt.Println("Classrooms:", d.Classrooms)
	fmt.Println("Questions length:", len(d.Questions))
*/
