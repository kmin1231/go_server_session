package service

import (
	"context"
	"fmt"

	"github.com/kmin1231/go_server_session/week07/Chapter19/entity"
	"github.com/kmin1231/go_server_session/week07/Chapter19/store"
)

type ListTask struct {
	DB   store.Queryer
	Repo TaskLister
}

func (l *ListTask) ListTasks(ctx context.Context) (entity.Tasks, error) {
	ts, err := l.Repo.ListTasks(ctx, l.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return ts, nil
}
