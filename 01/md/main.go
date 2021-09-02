package main

import "os"

func main() {
	createMarkDown()
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func createMarkDown(fileName ...string) {
	var filename string
	if len(fileName) == 0 {
		filename = "ReadMe.md"
	} else {
		filename = fileName[0]
	}

	f, err := os.Create(filename)
	defer f.Close()
	checkError(err)
	_, err = f.WriteString("## Helo MarkDown\n")
	checkError(err)
	_, err = f.WriteString("### Introduction\n")
	f.Sync()
}
