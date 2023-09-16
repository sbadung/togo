package todo

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID        uint64     `json: "id" gorm:"primary_gey"`
	Title     string     `json: "title"`
	Done      bool       `json: "done"`
	CreatedAt *time.Time `json: "created_at"`
}

func NewTodo(ID uint64, Title string, Done bool, CreatedAt *time.Time) *Todo {
	return &Todo{
		ID:        ID,
		Title:     Title,
		Done:      Done,
		CreatedAt: CreatedAt,
	}
}
