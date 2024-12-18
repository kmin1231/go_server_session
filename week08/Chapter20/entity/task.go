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
	ID			TaskID		`json:"id" db:"id"`
	UserID		UserID		`json:"user_id" db:"user_id"`
	Title   	string		`json:"title" db:"title"`
	Status  	TaskStatus	`json:"status" db:"status"`
	Created 	time.Time	`json:"created" db:"created"`
	Modified	time.Time	`json:"modified" db:"modified"`
}

// defines new type 'Tasks'
type Tasks []*Task
