package api

import (
	"log"

	iris "gopkg.in/kataras/iris.v6"

	"bitbucket.org/cdncache/taskmanager/svc"
	"bitbucket.org/cdncache/taskmanager/types"
	"bitbucket.org/cdncache/taskmanager/util"
)

// GetTask handles GET /tasks/:uuid
func GetTask(ctx *iris.Context) {}

// GetTasks handles GET /tasks/
// URL params:
//	limit:
//	offset:
//	from:
//	to:
func GetTasks(ctx *iris.Context) {}

// ModifyTask handles PUT /tasks/:uuid
func ModifyTask(ctx *iris.Context) {}

// StartTask handles PUT /tasks/:uuid/start
func StartTask(ctx *iris.Context) {}

// StopTask handles PUT /tasks/:uuid/stop
func StopTask(ctx *iris.Context) {
	resp := types.RespStopTask{}

	uuid := ctx.GetString("uuid")

	if len(uuid) == 0 {
		resp.Errmsg = "uuid not set in URL path"
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}

	err := svc.StopTask(uuid)
	if err != nil {
		resp.Errmsg = "stop task error: " + err.Error()
		if err == svc.ErrTaskNotFound {
			ctx.JSON(iris.StatusNotFound, resp)
			return
		}
		ctx.JSON(iris.StatusInternalServerError, resp)
		return
	}

	resp.Success = true
	ctx.JSON(iris.StatusOK, resp)
	return
}

// CreateTask handles POST /tasks/
func CreateTask(ctx *iris.Context) {
	req := types.ReqPostTask{}
	resp := types.RespPostTask{}

	err := ctx.ReadJSON(&req)
	if err != nil {
		log.Printf("parse request body to task error: %v\n", err)
		resp.Errmsg = "requst body unmarshal error: " + err.Error()
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}

	util.PrintPretty(req)

	if len(req.UserID) == 0 {
		resp.Errmsg = "user_id not set in request body"
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}
	if len(req.Name) == 0 {
		resp.Errmsg = "name not set in request body"
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}
	if len(req.LogDir) == 0 {
		resp.Errmsg = "log_dir not set in request body"
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}
	if req.CronInterval < 0 {
		resp.Errmsg = "cron_interval invalid in request body"
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}

	task, err := svc.CreateTask(req.UserID, req.Name, req.LogDir, req.CronInterval)
	if err != nil {
		log.Printf("create task error: %v\n", err)
		resp.Errmsg = "create task error: " + err.Error()
		ctx.JSON(iris.StatusInternalServerError, resp)
		return
	}

	resp.Success = true
	resp.Task = task
	ctx.JSON(iris.StatusOK, resp)
	return
}

// DeleteTask handles DELETE /tasks/:uuid
func DeleteTask(ctx *iris.Context) {
	resp := types.RespDeleteTask{}

	uuid := ctx.GetString("uuid")

	if len(uuid) == 0 {
		resp.Errmsg = "uuid not set in URL path"
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}

	err := svc.DeleteTask(uuid)
	if err != nil {
		resp.Errmsg = "delete task error: " + err.Error()
		if err == svc.ErrTaskNotFound {
			ctx.JSON(iris.StatusNotFound, resp)
			return
		}
		ctx.JSON(iris.StatusInternalServerError, resp)
		return
	}

	resp.Success = true
	ctx.JSON(iris.StatusOK, resp)
	return
}
