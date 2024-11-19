package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kmin1231/go_server_session/week07/Chapter19/entity"
	// "github.com/kmin1231/go_server_session/week07/Chapter19/store"
	"github.com/kmin1231/go_server_session/week07/Chapter19/testutil"
)

func TestListTask(t *testing.T) {
	type want struct {
		status  int
		rspFile string
	}
	tests := map[string]struct {
		// tasks map[entity.TaskID]*entity.Task
		tasks []*entity.Task
		want  want
	}{
		"ok": {
			// tasks: map[entity.TaskID]*entity.Task{
			tasks []*entity.Task{
				{
					ID:     1,
					Title:  "test1",
					Status: entity.TaskStatusTodo,
				},
				{
					ID:     2,
					Title:  "test2",
					Status: entity.TaskStatusDone,
				},
			},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_task/ok_rsp.json.golden",
			},
		},
		"empty": {
			// tasks: map[entity.TaskID]*entity.Task{},
			tasks: []*entity.Task{},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_task/empty_rsp.json.golden",
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/tasks", nil)

			moq := &ListTaskServiceMock{}
			moq.ListTtasksFunc = func(ctx context.Context) (entity.Tasks, error) {
				if tt.tasks != nil {
					return tt.tasks, nil
				}
				return nil, errors.New("error from mock")
			}
			// sut := ListTask{Store: &store.TaskStore{Tasks: tt.tasks}}
			sut := ListTask{Service: moq}
			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t,
				resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile),
			)
		})
	}
}
