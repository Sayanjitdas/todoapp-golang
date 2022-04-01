package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
	"todoapp/db"
	"todoapp/middleware"
)

//globals
var tpl *template.Template
var mux = http.NewServeMux()

const ADDR = "0.0.0.0:8000"

func TodoApp(w http.ResponseWriter, r *http.Request) {

	//POST
	if r.Method == http.MethodPost {
		todo := r.FormValue("todo")
		if len(todo) > 0 {
			//will be saving to db
			db.SaveTodos(todo)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}

	//DELETE
	if r.Method == http.MethodDelete {
		todoDeleteStruct := struct {
			TodoID int
		}{}
		err := json.NewDecoder(r.Body).Decode(&todoDeleteStruct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Panicln(err)
		}
		db.DeleteTodos(todoDeleteStruct.TodoID)
		json.NewEncoder(w).Encode(struct{ Msg string }{Msg: "success"})
		return
	}

	//GET
	type rowStruct struct {
		Id       int
		Todo     string
		Datetime string
	}
	data := []rowStruct{}
	result := db.ShowTodos()
	for result.Next() {
		row := rowStruct{}
		result.Scan(&row.Id, &row.Todo, &row.Datetime)
		data = append(data, row)
	}
	tpl.ExecuteTemplate(w, "index.html", data)
}

func init() {
	db.CreateTable()
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	tpl = template.Must(template.ParseGlob("./templates/*.html"))

	//Router Handler
	mux.HandleFunc("/", TodoApp)

}

func main() {
	wrappedMux := middleware.LoggerMiddleware(mux)
	log.Println("Listening on server 0.0.0.0:8000")
	http.ListenAndServe(ADDR, wrappedMux)
}
