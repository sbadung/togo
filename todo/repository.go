package todo

import "context"

type TodoRepository interface {
	FindAll(ctx context.Context) ([]Todo, error)
	Find(ctx context.Context, id uint64) (Todo, error)
	Insert(ctx context.Context, todo Todo) error
	Update(ctx context.Context, Todo Todo) error
	Delete(ctx context.Context, id uint64) error
	DeleteAll(ctx context.Context) error
}
