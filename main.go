package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const MAIN_URL = "http://localhost:3002"

const (
	ERR_CLI_SERVER        = "0"
	ERR_CLI_FORM_TITLE    = "1"
	ERR_CLI_FORM_USERNAME = "2"
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
	title := req.FormValue("title")
	username := req.FormValue("username")
	content := req.FormValue("content")

	// script 제거
	r := regexp.MustCompile(`<.*?script.*\/?>`)
	content = r.ReplaceAllString(content, "")

	if len(title) > 255 || len(title) == 0 {
		errCookie := http.Cookie{
			Name:     "errorServer",
			Value:    ERR_CLI_FORM_TITLE,
			SameSite: http.SameSiteLaxMode,
		}
		res.Header().Set("Set-Cookie", errCookie.String())
		http.Redirect(res, req, MAIN_URL, http.StatusSeeOther)

	} else if len(username) > 10 || len(username) < 1 {
		errCookie := http.Cookie{
			Name:     "errorServer",
			Value:    ERR_CLI_FORM_USERNAME,
			SameSite: http.SameSiteLaxMode,
		}
		res.Header().Set("Set-Cookie", errCookie.String())
		http.Redirect(res, req, MAIN_URL, http.StatusSeeOther)

	} else {

		// SQL 구현부

		http.Redirect(res, req, MAIN_URL, http.StatusSeeOther)

	}
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
