package fileserver

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadHandler(rw http.ResponseWriter, r *http.Request) {
	uploadFile, header, err := r.FormFile("upload_file")
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, err)
		return
	}
	defer uploadFile.Close()

	dirname := "./uploads"
	os.MkdirAll(dirname, 0777)
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(rw, err)
		return
	}
	io.Copy(file, uploadFile)
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, filepath)
}

func Start() {
	fmt.Println("File Server Started!")
	http.Handle("/", http.FileServer(http.Dir("fileserver/public")))

	http.HandleFunc("/uploads", uploadHandler)

	http.ListenAndServe(":3000", nil)
}