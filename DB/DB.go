package DB

import (
	"database/sql"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

var Database, errDB = sql.Open("sqlite3", "./products.db")

func OpenDB() error {
	slog.Info("Opening DB")

	if errDB != nil {
		return errDB
	}

	statement, errStatement := Database.Prepare("CREATE TABLE IF NOT EXISTS products (id INTEGER PRIMARY KEY, name TEXT, description TEXT, price REAL, category TEXT, created_at TEXT)")
	if errStatement != nil {
		return errStatement
	}
	statement.Exec()
	return nil
}

func CleanupDB() {
	Database.Close()
}
