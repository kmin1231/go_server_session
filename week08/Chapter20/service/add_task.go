package service

import (
	"context"
	"fmt"

	"github.com/kmin1231/go_server_session/week08/Chapter20/entity"
	"github.com/kmin1231/go_server_session/week08/Chapter20/store"
	"github.com/kmin1231/go_server_session/week08/Chapter20/auth"

)

type AddTask struct {
	DB   store.Execer
	Repo TaskAdder
}

func (a *AddTask) AddTask(ctx context.Context, title string) (*entity.Task, error) {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}
	
	t := &entity.Task{
		Title:  title,
		Status: entity.TaskStatusTodo,
	}
	err := a.Repo.AddTask(ctx, a.DB, t)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}
	return t, nil
}
