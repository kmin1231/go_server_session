package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	// "github.com/jmoiron/sqlx"
	"github.com/kmin1231/go_server_session/week07/Chapter19/entity"
	// "github.com/kmin1231/go_server_session/week07/Chapter19/store"
)

type AddTask struct {
	// DB        *sqlx.DB
	// Repo      *store.Repository
	Service   AddTaskService
	Validator *validator.Validate
}

func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		Title string `json:"title" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	if err := at.Validator.Struct(b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	// t := &entity.Task{
	// 	Title:  b.Title,
	// 	Status: entity.TaskStatusTodo,
	// }
	// err := at.Repo.AddTask(ctx, at.DB, t)

	t, err := at.Service.AddTask(ctx, b.Title)

	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := struct {
		ID entity.TaskID `json:"id"`
	}{ID: t.ID}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
