package types

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Task is the structure of table 'tasks' in database
// CronInterval: in second
// Status: CREATED / PENGDING / RUNNING / STOPPED / FAILED
type Task struct {
	gorm.Model
	UUID         string    `gorm:"primary_key column:user_id"`
	UserID       string    `gorm:"column:user_id"`
	Name         string    `gorm:"column:name"`
	LogDir       string    `gorm:"column:log_dir"`
	CronInterval int       `gorm:"column:cron_interval"`
	Status       string    `gorm:"column:status"`
	StartTime    time.Time `gorm:"column:start_time"`
	LastModified time.Time `gorm:"column:last_modified"`
}
