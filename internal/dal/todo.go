package dal

import (
	"context"
	"github.com/arnab-xyz/todoApp/internal/model"
	"gorm.io/gorm"
)

type Dal struct {
	db *gorm.DB
}

func NewDal(db *gorm.DB) *Dal {
	return &Dal{
		db: db,
	}
}

func (d *Dal) Create(ctx context.Context, todo *model.Todo) error {
	tx := d.db.Create(todo)
	return tx.Error
}

func (d *Dal) Get(ctx context.Context, id string) (*model.Todo, error) {
	var todo *model.Todo
	txn := d.db.First(&todo, "id = ?", id)
	if txn.Error != nil {
		return nil, txn.Error
	}
	return todo, nil
}

func (d *Dal) GetAll(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo
	txn := d.db.Find(&todos)
	if txn.Error != nil {
		return nil, txn.Error
	}
	return todos, nil
}

func (d *Dal) Update(ctx context.Context, id string, updatedTodo *model.Todo) error {
	txn := d.db.Save(updatedTodo)
	return txn.Error
}

func (d *Dal) Delete(ctx context.Context, id string) error {
	txn := d.db.Delete(&model.Todo{}, "id = ?", id)
	return txn.Error
}
