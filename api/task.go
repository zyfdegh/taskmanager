package api

import (
	"gopkg.in/kataras/iris.v6"
)

// GetTask handles GET /tasks/:id
func GetTask(ctx *iris.Context) {}

// GetTasks handles GET /tasks/
func GetTasks(ctx *iris.Context) {}

// ModifyTask handles PUT /tasks/:id
func ModifyTask(ctx *iris.Context) {}

// CreateTask handles POST /tasks/
func CreateTask(ctx *iris.Context) {}

// DeleteTask handles DELETE /tasks/:id
func DeleteTask(ctx *iris.Context) {}
