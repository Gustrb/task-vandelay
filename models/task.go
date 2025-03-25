package models

import "fmt"

type TaskStatus string

const (
	TaskStatusComplete  TaskStatus = "complete"
	TaskStatusToDo      TaskStatus = "todo"
	TaskStatusDoing     TaskStatus = "doing"
	TaskStatusCancelled TaskStatus = "cancelled"
)

func (ts *TaskStatus) Validate() error {
	switch *ts {
	case TaskStatusComplete, TaskStatusToDo, TaskStatusDoing, TaskStatusCancelled:
		return nil
	default:
		return fmt.Errorf("invalid task status %s", *ts)
	}
}

type Task struct {
	ID     string     `json:"id"`
	Value  string     `json:"value"`
	Status TaskStatus `json:"status"`
}

type CreateTaskPayload struct {
	Value string `json:"value"`
}
