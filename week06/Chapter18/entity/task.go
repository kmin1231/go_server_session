package entity

import "time"

// defines new types for 'TaskID' & 'TaskStatus'
type TaskID int64
type TaskStatus string

// defines constants for 'TaskStatus' type
const (
	TaskStatusTodo  TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone  TaskStatus = "done"
)

// defines'Task' struct with fields with JSON tag
type Task struct {
	ID      TaskID     `json:"id"`
	Title   string     `json:"title"`
	Status  TaskStatus `json:"status" `
	Created time.Time  `json:"created"`
}

// defines new type 'Tasks'
type Tasks []*Task
