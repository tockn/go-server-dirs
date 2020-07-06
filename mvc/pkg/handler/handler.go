package handler

import (
	"database/sql"
)

type Handler struct {
	db *sql.DB
}
