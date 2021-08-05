package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name 	  			 string
	Age 	  			 uint16
	Money     			 int16
	AvgGrades, Happiness float64
	Hobbies []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s." +
		" He is %d and he has money equel: %d",
		u.Name,
		u.Age,
		u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func homePage(w http.ResponseWriter, _ *http.Request)  {
	bob := User{
		Name: "Bob",
		Age: 19,
		Money: 100,
		AvgGrades: 7.77,
		Happiness: 0.777,
		Hobbies: []string{"Programming", "Sex"},
	}
	tmpl, _ := template.ParseFiles("wasm_game_of_life/www/index.html")
	err := tmpl.Execute(w, bob)
	if err != nil {
		return
	}
}

func contactsPage(w http.ResponseWriter, _ *http.Request)  {
	_, err := fmt.Fprintf(w, "Go and Rust the best programming languages! Contacts page!")
	if err != nil {
		return
	}
}

func handRequest()  {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func main()  {
	handRequest()
}
















