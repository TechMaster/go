package router

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"path"
)

type student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

/* Handle return plain text */
func returnPlainText(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home page"))
}

/* Handle return json */
func returnJson(w http.ResponseWriter, r *http.Request) {
	// Tạo mảng student
	students := []student{
		{Id: 1, Name: "Nguyễn Văn A"},
		{Id: 2, Name: "Trịnh Văn B"},
		{Id: 3, Name: "Ngô Thị C"},
	}

	studentsJson, err := json.Marshal(students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(studentsJson)
}

/* Handle return xml */
func returnXml(w http.ResponseWriter, r *http.Request) {
	students := []student{
		{Id: 1, Name: "Nguyễn Văn A"},
		{Id: 2, Name: "Trịnh Văn B"},
		{Id: 3, Name: "Ngô Thị C"},
	}

	studetsXml, err := xml.Marshal(students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(studetsXml)
}

func returnHtml(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta http-equiv="X-UA-Compatible" content="IE=edge">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Books</title>
			</head>
			<body>
				<h1>Những cuốn sách được yêu thích</h1>
				<ul>
					<li>Những người khốn khổ</li>
					<li>Đắc nhân tâm</li>
				</ul>
			</body>
		</html>
	`

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, html)
}

/* Handle return file */
func returnFile(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join("view", "index.html")
	// filePath := path.Join("static/image", "nature.jpg")

	http.ServeFile(w, r, filePath)
}

/* Handle health check server */
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("OK"))
}

func DemoRouter() {
	// Router
	http.HandleFunc("/", returnPlainText)
	http.HandleFunc("/json", returnJson)
	http.HandleFunc("/xml", returnXml)
	http.HandleFunc("/file", returnFile)
	http.HandleFunc("/html", returnHtml)
	http.HandleFunc("/health", HealthCheckHandler)

	// Serve static file
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Run server on port
	log.Fatal(http.ListenAndServe(":3000", nil))
}
