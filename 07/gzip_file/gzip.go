package gzip_file

import (
	"compress/gzip"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"
)

// Gzip Compression
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func Gzip(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzw := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		next.ServeHTTP(gzw, r)
	})
}

func handleWithoutGzip(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./view/index.html"))
	tmpl.Execute(w, nil)
}

func handleWithGzip(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./view/index.html"))
	tmpl.Execute(w, nil)
}

func DemoGzip() {
	mux := http.NewServeMux()

	handle := http.HandlerFunc(handleWithGzip)
	mux.Handle("/gzip", Gzip(handle))
	mux.HandleFunc("/", handleWithoutGzip)

	fmt.Println(http.ListenAndServe(":3000", mux))
}
