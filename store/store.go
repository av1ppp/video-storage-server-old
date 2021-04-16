package store

import (
	"database/sql"

	"github.com/AviParampampam/video-storage-server/config"
)

type Store struct {
	db  *sql.DB
	cfg *config.Config
}
