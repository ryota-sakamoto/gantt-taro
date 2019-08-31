package models

type Task struct {
	Id           int         `json:"id" db:"id"`
	Name         string      `json:"name" db:"name"`
	userId       int         `json:"user_id" db:"user_id"`
	projectId    int         `json:"project_id" db:"project_id"`
	parentTaskId int         `json:"parent_task_id" db:"parent_task_id"`
	StartedAt    timeWrapper `json:"started_at" db:"started_at"`
	EndedAt      timeWrapper `json:"ended_at" db:"ended_at"`
}

type timeWrapper []uint8

func (t timeWrapper) format() string {
	return string(t)
}

func (t timeWrapper) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.format() + `"`), nil
}
