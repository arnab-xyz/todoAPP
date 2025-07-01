package service

import (
	"context"
	"github.com/arnab-xyz/todoApp/internal/dal"
	"github.com/arnab-xyz/todoApp/internal/model"
	"github.com/google/uuid"
)

type Service struct {
	dal *dal.Dal
}

func NewService(dal *dal.Dal) *Service {
	return &Service{
		dal: dal,
	}
}

func (t *Service) Create(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	id := uuid.New().String()
	todo.ID = id
	err := t.dal.Create(ctx, todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *Service) Get(ctx context.Context, id string) (*model.Todo, error) {
	todo, err := t.dal.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *Service) GetAll(ctx context.Context) ([]*model.Todo, error) {
	todos, err := t.dal.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *Service) Update(ctx context.Context, id string, updatedTodo *model.Todo) error {
	updatedTodo.ID = id
	err := t.dal.Update(ctx, id, updatedTodo)
	return err
}

func (t *Service) Delete(ctx context.Context, id string) error {
	err := t.dal.Delete(ctx, id)
	return err
}
