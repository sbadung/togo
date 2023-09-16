package todo

import (
	"context"

	"gorm.io/gorm"
)

type PostgreSQLRepository struct {
	db *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) *PostgreSQLRepository {
	return &PostgreSQLRepository{db}
}

func (r *PostgreSQLRepository) FindAll(ctx context.Context) ([]Todo, error) {
	var todos []Todo
	if err := r.db.Find(&todos).Error; err != nil {
		return nil, ErrNotFound
	}
	return todos, nil
}

func (r *PostgreSQLRepository) Find(ctx context.Context, id uint64) (Todo, error) {
	var todo Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		return Todo{}, ErrNotFound
	}
	return todo, nil
}

func (r *PostgreSQLRepository) Insert(ctx context.Context, todo Todo) error {
	if err := r.db.Create(&todo).Error; err != nil {
		return ErrInsertFailed
	}
	return nil
}

func (r *PostgreSQLRepository) Update(ctx context.Context, todo Todo) error {
	if err := r.db.Save(&todo).Error; err != nil {
		return ErrUpdateFailed
	}
	return nil
}

func (r *PostgreSQLRepository) Delete(ctx context.Context, id uint64) error {
	if err := r.db.Delete(&Todo{}, id).Error; err != nil {
		return ErrDeleteAllFailed
	}
	return nil
}

func (r *PostgreSQLRepository) DeleteAll(ctx context.Context) error {
	if err := r.db.Delete(&Todo{}).Error; err != nil {
		return ErrDeleteAllFailed
	}
	return nil
}
