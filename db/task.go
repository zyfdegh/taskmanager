package db

import (
	"time"

	"bitbucket.org/cdncache/taskmanager/types"
)

func SaveTask(task types.Task) (savedTask types.Task, err error) {
	savedTask = task
	return
}

func UpdateTask(task types.Task) (updatedTask types.Task, err error) {
	updatedTask = task
	return
}

func DeleteTask(uuid string) (err error) {
	return
}

func GetTask(uuid string) (task types.Task, err error) {
	return
}

func QueryTasks(offset, limit int, from, to time.Time) (tasks []types.Task, err error) {
	return
}
