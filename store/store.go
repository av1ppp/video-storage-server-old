package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/AviParampampam/video-storage-server/config"
)

type Store struct {
	db  *sql.DB
	cfg *config.Config
}

func New(username, password, dbname string) (*Store, error) {
	s := &Store{}

	if err := s.configureDatabase(username, password, dbname, false); err != nil {
		return nil, err
	}

	return s, nil
}

func NewByConfig(cfg *config.Config) (*Store, error) {
	var (
		s        = &Store{cfg: cfg}
		username = cfg.Database.Username
		password = cfg.Database.Password
		dbname   = cfg.Database.DatabaseName
		sslmode  = cfg.Database.SSLMode
	)

	if err := s.configureDatabase(username, password, dbname, sslmode); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Store) configureDatabase(username, password, dbname string, sslmode bool) error {
	var sslmodeStr string
	if sslmode {
		sslmodeStr = "enable"
	} else {
		sslmodeStr = "disable"
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", username, password, dbname, sslmodeStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}
