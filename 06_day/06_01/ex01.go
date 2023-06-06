package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "mlarrauser" //"user-pg"
	password = "mlarrapswd" //"userpswd"
	dbname   = "mlarradb"   //"notes_db"
)

func admin(w http.ResponseWriter, r *http.Request) {
	log.Printf("admin\n")
	if r.URL.Path != "/admin" {
		http.NotFound(w, r)
		return
	}
	// Инициализируем срез содержащий пути к двум файлам. Обратите внимание, что
	// файл home.page.tmpl должен быть *первым* файлом в срезе.
	files := []string{
		"./ui/html/admin.html",
		"./ui/html/base.layout.html",
		"./ui/html/futer.partial.html",
	}

	// Используем функцию template.ParseFiles() для чтения файла шаблона.
	// Если возникла ошибка, мы запишем детальное сообщение ошибки и
	// используя функцию http.Error() мы отправим пользователю
	// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// Обработчик главной страницы.
func home(w http.ResponseWriter, r *http.Request) {
	log.Printf("home\n")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Инициализируем срез содержащий пути к двум файлам. Обратите внимание, что
	// файл home.page.tmpl должен быть *первым* файлом в срезе.
	files := []string{
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/futer.partial.html",
	}

	// Используем функцию template.ParseFiles() для чтения файла шаблона.
	// Если возникла ошибка, мы запишем детальное сообщение ошибки и
	// используя функцию http.Error() мы отправим пользователю
	// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// Обработчик для отображения содержимого заметки.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	// log.Printf("showsnippet\n")
	// w.Write([]byte("Отображение заметки..."))
	var (
		note_name    string
		note_content string
	)
	conn, err := pgx.Connect(context.Background(), "postgres://mlarrauser:mlarrapswd@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	numPage := 1 // read from body
	for i := numPage*3 - 2; i <= numPage*3; i++ {
		err = conn.QueryRow(context.Background(), "SELECT note_name, note_content FROM notes WHERE note_id=$1", i).Scan(&note_name, &note_content)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(w, "name: %s\ncontent: %s\n*********\n\n", note_name, note_content)
	}
	// err = conn.QueryRow(context.Background(), "SELECT note_name, note_content FROM notes WHERE note_id=$1", 9).Scan(&note_name, &note_content)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Fprintf(w, "name: %s\ncontent: %s\n*********\n\n", note_name, note_content)

}

// Обработчик для создания новой заметки.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	rt := r.URL.Query()
	if strings.EqualFold(rt.Get("name"), "Anton") && strings.EqualFold(rt.Get("password"), "12345") {

		log.Printf("/snippet/create\n")
		// w.Write([]byte("Форма для создания новой заметки..."))
		// log.Printf("admin\n")
		if r.URL.Path != "/snippet/create" {
			http.NotFound(w, r)
			return
		}

		// Инициализируем срез содержащий пути к двум файлам. Обратите внимание, что
		// файл home.page.tmpl должен быть *первым* файлом в срезе.
		files := []string{
			"./ui/html/create.page.html",
			"./ui/html/base.layout.html",
			"./ui/html/futer.partial.html",
		}

		// Используем функцию template.ParseFiles() для чтения файла шаблона.
		// Если возникла ошибка, мы запишем детальное сообщение ошибки и
		// используя функцию http.Error() мы отправим пользователю
		// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)
		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}

		// Затем мы используем метод Execute() для записи содержимого
		// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
		// возможность отправки динамических данных в шаблон.
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
			"./ui/html/autherror.html",
			"./ui/html/base.layout.html",
			"./ui/html/futer.partial.html",
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

func dbWriter(w http.ResponseWriter, r *http.Request) {

	note_name := ""
	note_content := ""

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), "postgres://mlarrauser:mlarrapswd@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// defer conn.Close(context.Background())

	if note_name != "" && note_content != "" {
		_, err = conn.Exec(context.Background(), "insert into notes(note_name, note_content) values($1, $2)", note_name, note_content)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Exec failed: %v\n", err)
			os.Exit(1)
		}
	}
	conn.Close(context.Background())
}

func main() {
	//	mux := http.NewServeMux()
	http.HandleFunc("/", home)
	http.HandleFunc("/admin", admin)
	// mux.HandleFunc("/js", downloadJs)
	// mux.HandleFunc("/css", downloadCss)
	// mux.HandleFunc("/изображения", downoadImages)
	http.HandleFunc("/snippet", showSnippet)
	http.HandleFunc("/snippet/create", createSnippet)
	http.HandleFunc("/snippet/write", dbWriter)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Используем функцию mux.Handle() для регистрации обработчика для
	// всех запросов, которые начинаются с "/static/". Мы убираем
	// префикс "/static" перед тем как запрос достигнет http.FileServer
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Go Backend: { HTTPVersion = 1 }; serving on http://localhost:8888/")
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}
