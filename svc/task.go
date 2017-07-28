package svc

import (
	"errors"
	"log"
	"time"

	"bitbucket.org/cdncache/taskmanager/db"
	"bitbucket.org/cdncache/taskmanager/types"
	"bitbucket.org/cdncache/taskmanager/util"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

func CreateTask(userID, taskName, logDir string, cronInterval int) (created *types.Task, err error) {
	task := types.Task{
		UserID:       userID,
		Name:         taskName,
		LogDir:       logDir,
		CronInterval: cronInterval,
	}

	task.UUID = util.NewUUID()
	task.Status = types.TaskStatusCreated

	now := time.Now()
	task.CreateAt = now
	task.UpdateAt = now

	util.PrintPretty(task)

	// save to db
	savedTask, err := db.SaveTask(task)
	if err != nil {
		log.Printf("save task to db error: %v\n", err)
		return
	}

	created = &savedTask
	return
}

func StartTask(uuid string) (err error) {
	task, err := db.GetTask(uuid)
	if err != nil {
		log.Printf("db get task error: %v\n", err)
		return
	}

	task.Status = types.TaskStatusPending
	task.UpdateAt = time.Now()

	_, err = db.UpdateTask(task)
	if err != nil {
		log.Printf("db update task error: %v\n", err)
		return
	}

	return
}

func StopTask(uuid string) (err error) {
	task, err := db.GetTask(uuid)
	if err != nil {
		log.Printf("db get task error: %v\n", err)
		return
	}

	task.Status = types.TaskStatusStopped
	task.UpdateAt = time.Now()

	_, err = db.UpdateTask(task)
	if err != nil {
		log.Printf("db update task error: %v\n", err)
		return
	}

	return
}

func GetTask(uuid string) (task *types.Task, err error) {
	dbTask, err := db.GetTask(uuid)
	if err != nil {
		log.Printf("db get task error: %v\n", err)
		return
	}

	task = &dbTask
	return
}

func GetTasks(limit, offset int, from, to time.Time) (tasks []types.Task, err error) {
	dbTasks, err := db.QueryTasks(limit, offset, from, to)
	if err != nil {
		log.Printf("query tasks error: %v\n", err)
		return
	}
	tasks = dbTasks
	return
}

func UpdateTask(task types.Task) (updated *types.Task, err error) {
	task.UpdateAt = time.Now()
	dbTask, err := db.UpdateTask(task)
	if err != nil {
		log.Printf("update task error: %v\n", err)
		return
	}

	updated = &dbTask
	return
}

func DeleteTask(uuid string) (err error) {
	err = StopTask(uuid)
	if err != nil {
		return
	}

	err = db.DeleteTask(uuid)
	if err != nil {
		log.Printf("delete task error: %v\n", err)
		return
	}
	return
}
