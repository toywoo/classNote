package main

import (
	"io/ioutil"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	htmlFile, readError := ioutil.ReadFile("./Client/note.html")
	if readError != nil {
		http.NotFound(w, req)
	} else {
		w.Write(htmlFile)
	}
}

func main() {
	http.Handle("/Client/", http.StripPrefix("/Client", http.FileServer(http.Dir("./Client"))))
	http.HandleFunc("/", index)
	http.ListenAndServe(":3002", nil)
}
