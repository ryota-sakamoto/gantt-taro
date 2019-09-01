package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Task struct {
	Id           int         `json:"id" db:"id"`
	Name         string      `json:"name" db:"name"`
	userId       int         `json:"user_id" db:"user_id"`
	projectId    int         `json:"project_id" db:"project_id"`
	parentTaskId int         `json:"parent_task_id" db:"parent_task_id"`
	StartedAt    timeWrapper `json:"started_at" db:"started_at"`
	EndedAt      timeWrapper `json:"ended_at" db:"ended_at"`
}

type timeWrapper struct {
	year  int
	month int
	day   int
}

func (t *timeWrapper) format() string {
	if t.year == 0 || t.month == 0 || t.day == 0 {
		return ""
	}
	return fmt.Sprintf("%04d-%02d-%02d", t.year, t.month, t.day)
}

func (t *timeWrapper) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.format() + `"`), nil
}

func (t *timeWrapper) UnmarshalJSON(b []byte) error {
	d := b[1 : len(b)-1]
	if len(d) != 10 {
		return nil
	}

	t.year, _ = strconv.Atoi(string(d[0:4]))
	t.month, _ = strconv.Atoi(string(d[5:7]))
	t.day, _ = strconv.Atoi(string(d[8:10]))

	return nil
}

func (t timeWrapper) ToTime() time.Time {
	d, _ := time.Parse("2006-01-02", t.format())
	return d
}

func (t *timeWrapper) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		t.year = v.Year()
		t.month = int(v.Month())
		t.day = v.Day()

		return nil
	}

	return errors.New("timeWrapper scan Failed")
}
