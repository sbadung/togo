package todo

import "errors"

var (
	ErrDatabaseInitialization = errors.New("Errore durante l'inizalizzazione del database")

	ErrNotFound        = errors.New("Todo not found")
	ErrInsertFailed    = errors.New("Failed to insert todo")
	ErrUpdateFailed    = errors.New("Failed to update todo")
	ErrDeleteFailed    = errors.New("Failed to delete todo")
	ErrDeleteAllFailed = errors.New("Failed to delete all todos")
	ErrInvalidId       = errors.New("Invalid ID")
	ErrInvalidBody     = errors.New("Invalid todo request body")
)
