package types

import (
	"time"
)

const (
	TaskStatusCreated = "CREATED"
	TaskStatusPending = "PENGDING"
	TaskStatusRunning = "RUNNING"
	TaskStatusStopped = "STOPPED"
	TaskStatusFailed  = "FAILED"
)

// Task is the structure of table 'tasks' in database
// CronInterval: in second, 0 means start immediately
// Status: CREATED / PENGDING / RUNNING / STOPPED / FAILED
type Task struct {
	UUID         string    `json:"uuid" gorm:"primary_key column:uuid"`
	UserID       string    `json:"user_id" gorm:"column:user_id"`
	Name         string    `json:"name" gorm:"column:name"`
	LogDir       string    `json:"log_dir" gorm:"column:log_dir"`
	CronInterval int       `json:"cron_interval" gorm:"column:cron_interval"`
	Status       string    `json:"status" gorm:"column:status"`
	StartAt      time.Time `json:"start_at" gorm:"column:start_at"`
	StopAt       time.Time `json:"stop_at" gorm:"column:stop_at"`
	CreateAt     time.Time `json:"create_at" gorm:"column:create_at"`
	UpdateAt     time.Time `json:"update_at" gorm:"column:update_at"`
}
