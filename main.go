package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func index(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "text/html; charset=utf-8")

	htmlFile, readError := ioutil.ReadFile("./Client/note.html")
	if readError != nil {
		http.NotFound(res, req)
	} else {
		res.Write(htmlFile)
	}
}

func save(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("save!"))
}

func handler(res http.ResponseWriter, req *http.Request) {
	var path = req.URL.Path
	splitedPath := strings.Split(path, "/")

	switch splitedPath[1] {
	case "":
		index(res, req)
	case "save":
		save(res, req)

	default:
		http.NotFound(res, req)
	}

}

func main() {
	http.Handle("/Client/", http.StripPrefix("/Client", http.FileServer(http.Dir("./Client"))))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3002", nil)
}
