package go_template

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func showTodo(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./view/todo.html"))
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl.Execute(w, data)
}

type Person struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func showPerson(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./view/person.html"))

	data := Person{
		Name:   "Nguyễn Văn A",
		Email:  "nguyenvana@gmail.com",
		Age:    35,
		Status: true,
	}
	tmpl.Execute(w, data)
}

func DemoTemplate() {
	http.HandleFunc("/", showTodo)
	http.HandleFunc("/person", showPerson)

	// Run server on port
	log.Fatal(http.ListenAndServe(":3000", nil))
}
