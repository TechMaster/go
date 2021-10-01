package querystring

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func userHandle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name, ok := query["name"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing name"))
		return
	}

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
