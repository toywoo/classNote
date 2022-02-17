package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	service "classNote/service"
)

type Handler struct {
	db *service.DB
}

const MAIN_URL = "http://localhost:3002"

const (
	ERR_CLI_SERVER        = "0"
	ERR_CLI_FORM_TITLE    = "1"
	ERR_CLI_FORM_USERNAME = "2"
)

// index
func index(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "text/html; charset=utf-8")

	htmlFile, readError := ioutil.ReadFile("./Client/note.html")
	if readError != nil {
		http.NotFound(res, req)
	} else {
		res.Write(htmlFile)
	}
}

// save
func save(res http.ResponseWriter, req *http.Request, db *service.DB) {
	title := req.FormValue("title")
	username := req.FormValue("username")
	content := req.FormValue("content")
	contentId := req.FormValue("id")

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
		isExistContent := db.IsExistContent(title, username, contentId)

		if isExistContent == service.INSERT {
			_, errInsertUser := db.Connection.Exec(`INSERT INTO note (username, title, content) VALUES ($1, $2, $3)`, username, title, content)

			if errInsertUser != nil {
				errCookie := http.Cookie{
					Name:     "errorServer",
					Value:    ERR_CLI_SERVER,
					SameSite: http.SameSiteLaxMode,
				}
				res.Header().Set("Set-Cookie", errCookie.String())
				http.Redirect(res, req, MAIN_URL, http.StatusSeeOther)
			}

			http.Redirect(res, req, MAIN_URL, http.StatusSeeOther)
		} else if isExistContent == service.UPDATE {
			_, errUpdateUser := db.Connection.Exec(`UPDATE note SET content=$1 WHERE username=$2 AND title=$3 AND id=$4`, content, username, title, contentId)

			if errUpdateUser != nil {
				errCookie := http.Cookie{
					Name:     "errorServer",
					Value:    ERR_CLI_SERVER,
					SameSite: http.SameSiteLaxMode,
				}
				res.Header().Set("Set-Cookie", errCookie.String())
				http.Redirect(res, req, MAIN_URL, http.StatusSeeOther)
			}

			http.Redirect(res, req, MAIN_URL, http.StatusSeeOther)
		} else {
			errCookie := http.Cookie{
				Name:     "errorServer",
				Value:    ERR_CLI_SERVER,
				SameSite: http.SameSiteLaxMode,
			}
			res.Header().Set("Set-Cookie", errCookie.String())
			http.Redirect(res, req, MAIN_URL, http.StatusSeeOther)
		}

	}
}

// get
// 		nav
func getNavContent(rows *sql.Rows) *[]byte {
	type NavContents struct {
		IDS    []int
		TITLES []string
		LEN    int
	}

	var ids []int
	var titles []string
	var len int = 0

	var id int
	var title string

	for rows.Next() {
		errRow := rows.Scan(&id, &title)
		if errRow != nil {
			log.Println("Error: Get new Row")
		}

		len++
		ids = append(ids, id)
		titles = append(titles, title)
	}

	var navContents NavContents
	navContents.IDS = ids
	navContents.TITLES = titles
	navContents.LEN = len

	jsonBytes, errJson := json.Marshal(navContents)
	if errJson != nil {
		log.Println("Error: Marshaling json")
	}

	return &jsonBytes // change dangling pointer?
}

func nav(res http.ResponseWriter, req *http.Request, db *service.DB) {
	rows, errUsername := db.Connection.Query(`SELECT id, title FROM note`)
	if errUsername != nil {
		errCookie := http.Cookie{
			Name:     "errorServer",
			Value:    ERR_CLI_SERVER,
			SameSite: http.SameSiteLaxMode,
		}
		res.Header().Set("Set-Cookie", errCookie.String())
		http.Redirect(res, req, MAIN_URL, http.StatusSeeOther)
	}

	jsonBytes := getNavContent(rows)

	rows.Close()

	res.Header().Add("Content-Type", "application/json; charset=utf-8")
	res.Write(*jsonBytes)
}

//		content
func content(res http.ResponseWriter, req *http.Request, db *service.DB) { // /content?id=
	type ContentInfo struct {
		TITLE        string
		USERNAME     string
		CREATED_TIME string
		CONTENT      string
	}
	var isErrConvert bool = false
	urlQuery := req.URL.Query()
	queriedId := urlQuery.Get("id")
	queryId, errConvert := strconv.ParseInt(queriedId, 10, 64)

	if errConvert != nil {
		log.Println("Error: Fail str to int Becas, id is strange") // 이게 망가져서 ErrCookie를 내면 결국에 전체가 망가지는 것이므로 아래에 쿠키 전달
		isErrConvert = true
	}

	var note ContentInfo

	db.Connection.QueryRow(`SELECT title, username, created_time, content FROM public.note WHERE id=$1`, queryId).Scan(&note.TITLE, &note.USERNAME, &note.CREATED_TIME, &note.CONTENT)

	jsonBytes, err := json.Marshal(note)
	if err != nil || isErrConvert == true {
		errCookie := http.Cookie{
			Name:     "errorServer",
			Value:    ERR_CLI_SERVER,
			SameSite: http.SameSiteLaxMode,
		}
		res.Header().Set("Set-Cookie", errCookie.String())
		http.Redirect(res, req, MAIN_URL, http.StatusSeeOther)
	} else {
		res.Header().Add("Content-Type", "application/json; charset=utf-8")
		res.Write(jsonBytes)
	}
}

func get(res http.ResponseWriter, req *http.Request, db *service.DB, splitedPath []string) {
	switch splitedPath[2] {
	case "nav":
		nav(res, req, db)
	case "content":
		content(res, req, db)
	default:
		http.NotFound(res, req)
	}
}

func (handler *Handler) pathNav(res http.ResponseWriter, req *http.Request) {
	var path = req.URL.Path
	splitedPath := strings.Split(path, "/")

	switch splitedPath[1] {
	case "":
		index(res, req)
	case "save":
		save(res, req, handler.db)

	case "get":
		get(res, req, handler.db, splitedPath)

	default:
		http.NotFound(res, req)
	}

}

func main() {
	handler := Handler{}

	db := service.NewDBConnection()
	handler.db = db

	http.Handle("/Client/", http.StripPrefix("/Client", http.FileServer(http.Dir("./Client"))))
	http.HandleFunc("/", handler.pathNav)
	http.ListenAndServe(":3002", nil)
}
