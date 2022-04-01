package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		log.Panicln(err.Error())
	}
}

func connect() *sql.DB {
	dbconn, err := sql.Open("sqlite3", "file:tododb.sql?_mutex=full")
	checkErr(err)
	err = dbconn.Ping()
	checkErr(err)
	return dbconn
}

func CreateTable() {
	query := `CREATE TABLE IF NOT EXISTS todoTable(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		todo TEXT NOT NULL,
		created_at TEXT DEFAULT CURRENT_TIMESTAMP
	);`
	db := connect()
	stmt, err := db.Prepare(query)
	checkErr(err)
	_, err = stmt.Exec()
	checkErr(err)
}

func SaveTodos(todo string) {

	query := `INSERT INTO todoTable(todo)VALUES(?);`

	db := connect()
	stmt, err := db.Prepare(query)
	checkErr(err)
	_, err = stmt.Exec(todo)
	checkErr(err)
}

func ShowTodos() *sql.Rows {

	query := `SELECT id,todo,datetime(created_at,'localtime') as created_at FROM todoTable ORDER BY created_at desc;`
	db := connect()
	rows, err := db.Query(query)
	checkErr(err)
	return rows
}

func DeleteTodos(todoId int) {
	query := `DELETE FROM todoTable WHERE id=?;`
	db := connect()
	stmt, err := db.Prepare(query)
	checkErr(err)
	_, err = stmt.Exec(todoId)
	checkErr(err)
}
