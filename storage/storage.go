package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

//Instance
type Storage struct {
	config *Config
	//Database file descriptor
	db *sql.DB
}

//Storage Constructor
func New(config *Config) *Storage {
	return &Storage{config: config}
}

//Open connection method
func (s *Storage) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	s.db = db
	log.Println("Database connection successfully")
	return nil
}

//Close connection
func (s *Storage) Close() {
	s.db.Close()
}
