package database

import "database/sql"

func CreateTables(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY,
            name TEXT
        )
    `)
	if err != nil {
		return err
	}

	// Add more table creation statements here if necessary

	return nil
}
