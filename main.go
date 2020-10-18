package sqliteconnection

import (
	"database/sql"
	"errors"
	"strings"
)

// New returns new DB connection that is checked already
func New(pathToDB string) (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", strings.Join([]string{pathToDB, "?_foreign_keys=on"}, ""))
	if err != nil {
		return
	}

	err = checkForeignKeys(db)
	return
}

// checkForeignKeys additionally checks if the connection is properly set up and established
func checkForeignKeys(db *sql.DB) error {
	query, err := db.Query(`PRAGMA foreign_keys;`)
	if err != nil {
		return err
	}
	defer query.Close()

	var result string
	ok := query.Next()
	if ok {
		err = query.Scan(&result)
		if err != nil {
			return err
		}
	}

	if result == "0" {
		return errors.New("incorrect 'foreign_keys' setting for the database")
	}
	return nil
}