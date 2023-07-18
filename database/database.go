package database

import (
	"database/sql"
	"log"
	"reflect"

	_ "github.com/lib/pq"
)

const tagName = "sql"

type Database interface {
	Close() error
	CreateTables() error
	DropTables() error
	Ping() error
	Query(query string, args ...any) (*sql.Rows, error)
}

type Db struct {
	database *sql.DB
}

func Connect(str string) (Database, error) {
	db, err := sql.Open("postgres", str)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return Wrap(db), nil
}

func Wrap(db *sql.DB) Database {
	return &Db{
		database: db,
	}
}

func (db *Db) Close() error {
	return db.database.Close()
}

func (db *Db) CreateTables() error {
	t := `CREATE TABLE IF NOT EXISTS groups (
		"id" TEXT NOT NULL PRIMARY KEY,
		"name" TEXT NOT NULL
	);`

	s, err := db.prepare(t)
	if err != nil {
		log.Fatal(err)
		return err
	}

	s.Exec()
	if _, err = s.Exec(); err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("groups table created successfully")

	t = `CREATE TABLE IF NOT EXISTS services (
		"id" TEXT NOT NULL PRIMARY KEY,
		"group_id" TEXT NOT NULL,
		"name" TEXT NOT NULL,
		"description" TEXT NULL,
		"is_enabled" BIT NOT NULL,
		"is_public" BIT NOT NULL,
		"domain" TEXT NOT NULL,
		"port" INT NOT NULL,
		"interval" INT NOT NULL,
		"timeout" INT NOT NULL,
		"type" TEXT NOT NULL,
		"method" TEXT NOT NULL,
		"follow_redirects" BIT NOT NULL,
		"verify_ssl" BIT NOT NULL,
		"created_at" BIGINT NOT NULL,
		"updated_at" BIGINT NULL,
		"last_check_at" BIGINT NULL,
		"allow_notifications" BIT NOT NULL,
		"notify_after" BIGINT NULL,
		"notify_all_changes" BIT NOT NULL,
		CONSTRAINT FK_GROUP_ID
			FOREIGN KEY (group_id)
				REFERENCES groups(id)
	);`

	s, err = db.prepare(t)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if _, err = s.Exec(); err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("services table created successfully")

	return nil
}

func (db *Db) DropTables() error {
	t := `DROP TABLE IF EXISTS services;`

	s, err := db.prepare(t)
	if err != nil {
		log.Fatal(err)
		return err
	}

	s.Exec()
	log.Println("Services table dropped successfully.")

	t = `DROP TABLE IF EXISTS groups;`

	s, err = db.prepare(t)
	if err != nil {
		log.Fatal(err)
		return err
	}

	s.Exec()
	log.Println("Groups table dropped successfully.")

	return nil
}

func (db *Db) Ping() error {
	return db.database.Ping()
}

func (db *Db) Query(query string, args ...any) (*sql.Rows, error) {
	return db.database.Query(query, args)
}

func (db *Db) Save(value interface{}) Database {
	v := reflect.ValueOf(value)

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get(tagName)
		if tag == "" || tag == "-" {
			continue
		}

		log.Printf("Found tag: %s", tag)
	}

	return db
}

func (db *Db) prepare(text string) (*sql.Stmt, error) {
	return db.database.Prepare(text)
}
