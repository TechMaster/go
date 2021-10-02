package querystring

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// http:localhost:3000
func userHandle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name, ok := query["name"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing name"))
		return
	}

	// ["a", "b", "c"] => a,b,c
	message := fmt.Sprintf("Hello %s", strings.Join(name, ","))
	w.Write([]byte(message))
}

func DemoQueryString() {
	http.HandleFunc("/", userHandle)
	fmt.Println(http.ListenAndServe(":3000", nil))
}

// ============= User ============
// {{.ID}} => Go template
var tplStr = `
<html>
	<h1>Customer {{.ID}}
	{{if .ID }}
	<p>Details:</p>
	<ul>
		{{if .Name}}<li>Name: {{.Name}}</li>{{end}}
		{{if .Age}}<li>Age: {{.Age}}</li>{{end}}
	</ul>
	{{else}}
	<p>Data not available</p>
	{{end}}
</html>
`

type Customer struct {
	ID   int
	Name string
	Age  int
}

func customerHandle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	customer := Customer{}

	id, ok := query["id"]
	if ok {
		customer.ID, _ = strconv.Atoi(strings.Join(id, ","))
	}

	name, ok := query["name"]
	if ok {
		customer.Name = strings.Join(name, ",")
	}

	age, ok := query["age"]
	if ok {
		customer.Age, _ = strconv.Atoi(strings.Join(age, ""))
	}

	tmpl, _ := template.New("test").Parse(tplStr)
	tmpl.Execute(w, customer)
}

func DemoQueryStringWithTemplate() {
	http.HandleFunc("/", customerHandle)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// ============= Book ============
type book struct {
	ID     int
	Name   string
	Author string
	Year   int
}

func bookHandle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	books := []book{
		{ID: 1, Name: "book1", Author: "author1", Year: 2020},
		{ID: 2, Name: "book2", Author: "author2", Year: 2021},
		{ID: 3, Name: "book3", Author: "author3", Year: 2019},
		{ID: 4, Name: "book4", Author: "author4", Year: 2020},
	}

	var booksReturn []book
	name, ok := query["name"]
	if ok {
		filterName := strings.Join(name, ",")
		booksReturn = filterByName(books, filterName)
	} else {
		booksReturn = append(booksReturn, books...)
	}

	year, ok := query["year"]
	if ok {
		filterYear, _ := strconv.Atoi(strings.Join(year, ""))
		booksReturn = filterByYear(booksReturn, filterYear)
	}

	booksJson, err := json.Marshal(booksReturn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(booksJson)
}

func filterByName(books []book, name string) (result []book) {
	for _, book := range books {
		if book.Name == name {
			result = append(result, book)
		}
	}
	return result
}

func filterByYear(books []book, year int) (result []book) {
	for _, book := range books {
		if book.Year == year {
			result = append(result, book)
		}
	}
	return result
}

func DemoQueryStringWithBook() {
	http.HandleFunc("/", bookHandle)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
