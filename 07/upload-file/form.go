package uploadfile

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	// Parse raw query từ url
	r.ParseForm()

	// In thông tin form
	fmt.Println(r.Form)

	// In thông tin từ path gửi request
	fmt.Println("path", r.URL.Path)

	// In các giá trị key-value từ query
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Jenny!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./view/login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		
		// In ra thông tin form
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func DemoForm() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	fmt.Println(http.ListenAndServe(":3000", nil))
}
