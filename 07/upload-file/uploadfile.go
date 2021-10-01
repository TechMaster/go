package uploadfile

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./view/upload.html")
		t.Execute(w, nil)
	} else {
		// Đọc dữ liệu từ form file
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		// Trả về kết quả cho client để hiển thị
		fmt.Fprintf(w, "%v", handler.Header)

		// Mở file. Trường hợp không tồn tại file thì tạo mới
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func DemoUploadFile() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":3000", nil)
}
