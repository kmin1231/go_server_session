package handler

import (
	"net/http"

	"github.com/kmin1231/go_server_session/week07/Chapter19/entity"
	// "github.com/kmin1231/go_server_session/week07/Chapter19/store"
	// "github.com/kmin1231/go_server_session/week07/Chapter19/store"
)

type ListTask struct {
	// DB   *sqlx.DB
	// Repo *store.Repository
	Service ListTasksService
}

// func (l *ListTask) ListTasks(ctx context.Context) (entity.Tasks, error) {
// 	ts, err := l.Repo.ListTasks(ctx, l.DB)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to list: %w", err)
// 	}
// 	return ts, nil
// }

type task struct {
	ID     entity.TaskID     `json:"id"`
	Title  string            `json:"title"`
	Status entity.TaskStatus `json:"status"`
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// tasks, err := lt.Repo.ListTasks(ctx, lt.DB)
	tasks, err := lt.Service.ListTasks(ctx)

	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := []task{}
	for _, t := range tasks {
		rsp = append(rsp, task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
