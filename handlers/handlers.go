package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"text/template"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func Prepare() {
	//http.HandleFunc("/", Index)
	http.HandleFunc("/session/login", login)
	http.HandleFunc("/session/signup", signup)
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.ListenAndServe(":8080", nil)
}
func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("session/signup.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("email:", r.Form["email"])
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		if data.isUnAvailable() {
			fmt.Errorf("The current username is unavailable")
		} else {
			t, _ := template.ParseFiles("session/login.html")
			t.Execute(w, nil)
		}

	}
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, err := template.ParseFiles("session/login.html")
		fmt.Println("yooooo", err)
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}
func index(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Method)
	//f, _ := file.Load("web/index.html")
	//fmt.Fprintf(w, string(f.Body), r.URL.Path[1:])
}
