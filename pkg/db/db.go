package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

// Params represents the project, user, and token parameters.
type Params struct {
	Domain  string
	Project string
	User    string
	Token   string
}

// NewDB creates a new DB instance with the specified database file path.
func NewDB() (*DB, error) {
	// Open an SQLite database file or create one if it doesn't exist
	db, err := sql.Open("sqlite3", "jira-sync.db")
	if err != nil {
		return nil, err
	}

	// Create the "params" table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS params (
		domain VARCHAR,
		project VARCHAR,
		user VARCHAR,
		token VARCHAR
	)`)
	if err != nil {
		return nil, err
	}

	return &DB{
		db: db,
	}, nil
}

func (d *DB) GetParams() (Params, error) {
	// d.db.Exec("UPDATE params SET project = ''")

	var params Params
	err := d.db.QueryRow("SELECT domain, project, user, token FROM params LIMIT 1").Scan(&params.Domain, &params.Project, &params.User, &params.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			return Params{}, nil
		}
		return Params{}, err
	}

	return params, nil
}

func (d *DB) SetParams(domain, project, user, token string) error {
	err := d.DeleteParams()
	if err != nil {
		return err
	}

	_, err = d.db.Exec("INSERT INTO params (domain, project, user, token) VALUES (?, ?, ?, ?)", domain, project, user, token)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) DeleteParams() error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM params")
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
