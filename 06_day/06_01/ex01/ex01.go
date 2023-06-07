package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/jackc/pgx/v5"
)

type sNote struct {
	Id      int
	Name    string
	Content string
	Hidden  string
}

type sSend struct {
	CountPage  int
	Page       int
	Id0        int
	Name0      string
	Content0   template.HTML
	Hidden0    string
	Id1        int
	Name1      string
	Content1   template.HTML
	Hidden1    string
	Id2        int
	Name2      string
	Content2   template.HTML
	Hidden2    string
	HiddenNext string
	HiddenPrev string
	PageNext   int
	PagePrev   int
}

type postSend struct {
	Page    int
	Id      int
	Name    string
	Content template.HTML
}

type sProgramm struct {
	Programms []sTypeProgramm `json:"programm"`
}

type sTypeProgramm struct {
	SnippetUsers  []sSnippetUser  `json:"snippet"`
	PostgresUsers []sPostgresUser `json:"postgres"`
}

type sSnippetUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type sPostgresUser struct {
	Name     string `json:"user_name"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}

func admin(w http.ResponseWriter, r *http.Request) {
	log.Printf("admin\n")
	if r.URL.Path != "/admin" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./html/admin.html",
		"./html/base.layout.html",
		"./html/futer.partial.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Printf("home\n")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	var numPage int
	if !(strings.EqualFold(r.URL.RawQuery, "")) {
		rQuery := r.URL.Query()
		s, _ := strconv.ParseInt(rQuery.Get("page"), 10, 16)
		numPage = int(s)
	}

	var countNotes int

	userData := &sProgramm{}
	err := userData.getUserData()
	if err != nil {
		log.Println("Bad users data\n")
		return
	}

	urlPostgres := "postgres://" + userData.Programms[1].PostgresUsers[0].Name + ":" + userData.Programms[1].PostgresUsers[0].Password + "@localhost:5432/" + userData.Programms[1].PostgresUsers[0].DBName
	conn, err := pgx.Connect(context.Background(), urlPostgres)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	err = conn.QueryRow(context.Background(), "SELECT COUNT(note_id) FROM notes ").Scan(&countNotes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	conn, err = pgx.Connect(context.Background(), urlPostgres)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	notesOnPage := 3
	queryText := `
	SELECT *
	FROM notes 
	ORDER BY note_id DESC
	LIMIT $1
	OFFSET $2`

	rows, err := conn.Query(context.Background(), queryText, notesOnPage, notesOnPage*numPage)
	if err != nil {
		fmt.Errorf("query notes failed: %v\n", err)
		return
	}
	defer rows.Close()
	var notes []sNote
	for rows.Next() {
		note := sNote{}
		err = rows.Scan(&note.Id, &note.Name, &note.Content)
		if err != nil {
			fmt.Errorf("query notes failed: %v\n", err)
			return
		}
		notes = append(notes, note)
	}

	if len(notes) != notesOnPage {
		lenght := len(notes)
		for i := 3; i > lenght; i-- {
			note := sNote{
				Id:      0,
				Hidden:  "hidden",
				Name:    "",
				Content: "",
			}
			notes = append(notes, note)
		}
	}

	// Инициализируем срез содержащий пути к двум файлам. Обратите внимание, что
	// файл home.page.tmpl должен быть *первым* файлом в срезе.
	files := []string{
		"./html/home.page.html",
		"./html/base.layout.html",
		"./html/futer.partial.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	dataForExec := sSend{
		Page:     numPage + 1,
		Id0:      notes[0].Id,
		Name0:    notes[0].Name,
		Content0: template.HTML(mdToHTML([]byte(notes[0].Content))),
		Hidden0:  notes[0].Hidden,
		Id1:      notes[1].Id,
		Name1:    notes[1].Name,
		Content1: template.HTML(mdToHTML([]byte(notes[1].Content))),
		Hidden1:  notes[1].Hidden,
		Id2:      notes[2].Id,
		Name2:    notes[2].Name,
		Content2: template.HTML(mdToHTML([]byte(notes[2].Content))),
		Hidden2:  notes[2].Hidden,
	}

	if numPage == 0 {
		dataForExec.HiddenPrev = "hidden"
	} else {
		dataForExec.PagePrev = numPage - 1
	}
	if countNotes > (numPage+1)*3 {
		dataForExec.PageNext = numPage + 1
	} else {
		dataForExec.HiddenNext = "hidden"
	}

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, dataForExec)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}

// Обработчик для отображения содержимого заметки.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	log.Printf("/snippet/showpost\n")
	if r.URL.Path != "/showpost" {
		http.NotFound(w, r)
		return
	}
	postData := postSend{}
	rQuery := r.URL.Query()
	ps, _ := strconv.ParseInt(rQuery.Get("page"), 10, 16)
	postData.Page = int(ps) - 1
	is, _ := strconv.ParseInt(rQuery.Get("id"), 10, 16)
	postData.Id = int(is)

	queryText := "SELECT * FROM notes WHERE note_id=$1"

	userData := &sProgramm{}
	err := userData.getUserData()
	if err != nil {
		log.Println("Bad users data\n")
		return
	}

	urlPostgres := "postgres://" + userData.Programms[1].PostgresUsers[0].Name + ":" + userData.Programms[1].PostgresUsers[0].Password + "@localhost:5432/" + userData.Programms[1].PostgresUsers[0].DBName
	conn, err := pgx.Connect(context.Background(), urlPostgres)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var tempStr string
	err = conn.QueryRow(context.Background(), queryText, &postData.Id).Scan(&postData.Id, &postData.Name, &tempStr)
	if err != nil {
		fmt.Errorf("query notes failed: %v\n", err)
		return
	}
	postData.Content = template.HTML(mdToHTML([]byte(tempStr)))
	log.Println(postData.Id, postData.Name, postData.Content)
	if err != nil {
		log.Println(postData.Id, postData.Name, postData.Content)
		fmt.Errorf("query notes failed: %v\n", err)
		return
	}

	files := []string{
		"./html/post.page.html",
		"./html/base.layout.html",
		"./html/futer.partial.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, postData)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// Парсер формы с заметкой
func parceSnippet(str string) (string, string) {
	i1 := strings.Index(str, "name=")
	i2 := strings.Index(str, "&snippet=")
	return str[i1+5 : i2], str[i2+9:]
}

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

// Обработчик записи заметки в БД
func writeSnippet(w http.ResponseWriter, r *http.Request) {
	log.Println("/snippet/write")
	body, _ := ioutil.ReadAll(r.Body)
	str, _ := url.PathUnescape(strings.Replace(string(body), "+", " ", -1))

	note_name, note_content := parceSnippet(str)
	note_contentHtml := mdToHTML([]byte(note_content))

	userData := &sProgramm{}
	err := userData.getUserData()
	if err != nil {
		log.Println("Bad users data\n")
		return
	}

	urlPostgres := "postgres://" + userData.Programms[1].PostgresUsers[0].Name + ":" + userData.Programms[1].PostgresUsers[0].Password + "@localhost:5432/" + userData.Programms[1].PostgresUsers[0].DBName
	conn, err := pgx.Connect(context.Background(), urlPostgres)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	if note_name != "" && note_content != "" {
		_, err := conn.Exec(context.Background(), "INSERT INTO notes(note_name, note_content) VALUES($1, $2)", note_name, note_contentHtml)
		if err != nil {
			log.Printf("Exec failed: %v\n", err)
			os.Exit(1)
		}
	}
	conn.Close(context.Background())

	log.Printf("/snippet/write\n")
	if r.URL.Path != "/snippet/write" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./html/writedDB.page.html",
		"./html/base.layout.html",
		"./html/futer.partial.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (p *sProgramm) getUserData() error {
	buff, err := os.ReadFile("dataUsers.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(buff, &p)
	return err
}

// Обработчик для создания новой заметки.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	rt := r.URL.Query()

	userData := &sProgramm{}
	err := userData.getUserData()
	if err != nil {
		log.Println("Bad users data\n")
		return
	}

	if strings.EqualFold(rt.Get("name"), userData.Programms[0].SnippetUsers[0].Name) && strings.EqualFold(rt.Get("password"), userData.Programms[0].SnippetUsers[0].Password) {

		log.Printf("/snippet/create\n")
		if r.URL.Path != "/snippet/create" {
			http.NotFound(w, r)
			return
		}

		files := []string{
			"./html/create.page.html",
			"./html/base.layout.html",
			"./html/futer.partial.html",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}

		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	} else {
		if r.URL.Path != "/snippet/create" {
			http.NotFound(w, r)
			return
		}

		files := []string{
			"./html/autherror.html",
			"./html/base.layout.html",
			"./html/futer.partial.html",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}

		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}
}
