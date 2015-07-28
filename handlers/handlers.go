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
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.Handle("/", http.FileServer(http.Dir("./html")))
	http.ListenAndServe(":8080", nil)
}
func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("html/indexOld.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("email:", r.Form["email"])
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("html/indexOld.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		http.Redirect(w, r, "html/indexOld.html", http.StatusFound)
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
	//f, _ := file.Load("html/index.html")
	//fmt.Fprintf(w, string(f.Body), r.URL.Path[1:])
}
