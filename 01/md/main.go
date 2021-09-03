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

/*
Tạo file Markdown, nếu không truyền vào fileName thì tạo file ReadMe.md
nếu truyền vào fileName thì tạo file fileName.md
*/
func createMarkDown(fileName ...string) {
	var filename string
	if len(fileName) == 0 {
		filename = "ReadMe.md"
	} else {
		filename = fileName[0]
	}

	f, err := os.Create(filename)
	checkError(err)
	defer f.Close() //hãy đóng file f khi hàm thoát

	_, err = f.WriteString("## Helo MarkDown\n")
	checkError(err)

	_, err = f.WriteString("### Introduction\n")
	checkError(err)
	err = f.Sync()
	checkError(err)
}
