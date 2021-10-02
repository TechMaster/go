package router

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestHealthCheckHandle(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/health")
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	sb := string(body)

	want := "OK"

	if sb != want {
		t.Errorf("handler returned unexpected body: got %v want %v", sb, want)
	}
}
