package main

import (
	"encoding/csv"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"log"
)

var sessionStore = sessions.NewCookieStore([]byte("03uofewöjcnvgfo32mv 9ur0ßur0ß9 "))

type Session struct {
	id string
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("dotfiles.html/root.html"))
	t.Execute(w, nil)
}

func handlePgpassFile(w http.ResponseWriter, r *http.Request) {

	sess, err := sessionStore.Get(r, "dotSimple")
	if err != nil {
		http.error(w, err.Error(), http.StatusInternalServerError)
		log.Print(err)
		return
	}
	defer sess.Save(r, w)

	username = r.BasicAuth()

	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", "inline; filename=\"_pgpass\"")
		csvWriter := csv.NewWriter(w)
		csvWriter.Comma = ':'
		csvWriter.Write([]string{"#host", "port", "database", "username", "password"})
		for _, host := range []string{"vie-bio-postgres", "vieciaepg"} {
			csvWriter.Write([]string{host, "port", "database", "username", "password"})
		}
		csvWriter.Flush()
	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

	}
}

func main() {

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/pgpass", handlePgpassFile)

	http.ListenAndServe("[::1]:8080", nil)
}
