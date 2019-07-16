package models

import "time"

type Task struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	UserId       int       `json:"user_id"`
	projectId    int       `json:"project_id"`
	parentTaskId int       `json:"parent_task_id"`
	StartedAt    time.Time `json:"started_at"`
	EndedAt      time.Time `json:"ended_at"`
}
