// main.go
package main

import (
	"github.com/website-golang/data"
)

func main() {
	//var user1 database.User
	user1 := data.CreateRandomUser() // return a *user
	user1.PrintUser()
	user2 := data.CreateRandomUser() // return a *user
	user2.PrintUser()
}
