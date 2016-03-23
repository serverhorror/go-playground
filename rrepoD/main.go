package main

import (
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"path"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received a %q reqeust", r.Method)
	switch {
	// case r.Method == "GET":
	case r.Method == "POST":
		r.ParseMultipartForm(int64(math.Pow(2, 20) * 100))
		file, header, err := r.FormFile("upload")
		if err != nil {
			log.Printf("err(http.Request.FormFile): %q", err)
			return
		}
		defer file.Close()
		log.Printf("header: %q", header.Header)
		tmp, err := ioutil.TempFile("", "")
		if err != nil {
			log.Printf("err(ioutil.TempFile): %q", err)
			return
		}
		log.Printf("tmp.Name(): %q", tmp.Name())
		defer os.Remove(tmp.Name())
		// defer tmp.Close()
		io.Copy(tmp, file)
		newName := path.Join("C:\\Users\\marcherm\\go\\src\\github.com\\serverhorror\\go-playground\\rrepoD", "cbdd2.tar.gz")
		log.Printf("newName: %q", newName)
		tmp.Close()
		err = os.Rename(tmp.Name(), newName)
		if err != nil {
			log.Printf("err(os.Rename): %q", err)
		}

	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

}

func main() {
	http.HandleFunc("/api/upload", uploadHandler)
	http.ListenAndServe("[::1]:8080", nil)
}
