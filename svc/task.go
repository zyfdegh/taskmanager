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

func CreateTask(userID, taskName, logDir string, cronInterval int) (createdTask *types.Task, err error) {
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

	createdTask = &savedTask
	return
}

func StopTask(uuid string) (err error) {
	return
}

func GetTask(uuid string) (task *types.Task, err error) {
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
