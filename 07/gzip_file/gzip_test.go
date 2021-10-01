package gzip_file

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func sendRequest(client *http.Client, addr string) {
	res, err := client.Get(addr)
	if err != nil {
		panic(err)
	}

	if res.StatusCode != 200 {
		panic("request failed")
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	err = res.Body.Close()
	if err != nil {
		panic(err)
	}
}

func Benchmark_NoGzip(b *testing.B) {
	client := &http.Client{}

	for n := 0; n < b.N; n++ {
		sendRequest(client, "http://localhost:3000")
	}
}

func Benchmark_Gzip(b *testing.B) {
	client := &http.Client{}

	for n := 0; n < b.N; n++ {
		sendRequest(client, "http://localhost:3000/gzip")
	}
}
