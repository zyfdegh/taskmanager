package api

import (
	"log"
	"time"

	iris "gopkg.in/kataras/iris.v6"

	"bitbucket.org/cdncache/taskmanager/svc"
	"bitbucket.org/cdncache/taskmanager/types"
	"bitbucket.org/cdncache/taskmanager/util"
)

const (
	UrlTimeLayout = time.RFC3339
)

// GetTask handles GET /tasks/:uuid
// ==== Request ====
//
// ==== Response ====
// Body: Task
func GetTask(ctx *iris.Context) {
	resp := types.RespGetTask{}

	uuid := ctx.GetString("uuid")
	if len(uuid) == 0 {
		resp.Errmsg = "'uuid' not set in URL path"
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}

	task, err := svc.GetTask(uuid)
	if err != nil {
		resp.Errmsg = "get task error: " + err.Error()
		if err == svc.ErrTaskNotFound {
			ctx.JSON(iris.StatusBadRequest, resp)
			return
		}
		ctx.JSON(iris.StatusInternalServerError, resp)
		return
	}

	resp.Success = true
	resp.Task = *task
	ctx.JSON(iris.StatusOK, resp)
	return
}

// GetTasks handles GET /tasks/
// ==== Request ====
// URL params:
//	limit: 10
//	offset: 20
//	from: <SOME_TIME>
//	to: <SOME_TIME>
//
// ==== Response ====
func GetTasks(ctx *iris.Context) {
	resp := types.RespGetTasks{}

	limit, err := ctx.URLParamInt("limit")
	if err != nil {
		resp.Errmsg = "parse url param 'limit' to int error: " + err.Error()
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}
	offset, err := ctx.URLParamInt("offset")
	if err != nil {
		resp.Errmsg = "parse url param 'offset' to int error: " + err.Error()
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}
	from, err := time.Parse(UrlTimeLayout, ctx.URLParam("from"))
	if err != nil {
		resp.Errmsg = "parse url param 'from' to time error: " + err.Error()
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}
	to, err := time.Parse(UrlTimeLayout, ctx.URLParam("to"))
	if err != nil {
		resp.Errmsg = "parse url param 'to' to time error: " + err.Error()
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}

	tasks, err := svc.GetTasks(limit, offset, from, to)
	if err != nil {
		log.Printf("get tasks error: %v\n", err)
		resp.Errmsg = "get tasks error: " + err.Error()
		ctx.JSON(iris.StatusInternalServerError, resp)
		return
	}

	resp.Success = true
	resp.Tasks = tasks
	return
}

// ModifyTask handles PUT /tasks/:uuid
// ==== Request ====
// Body: Task
//
// ==== Response ====
// Body: Task
func ModifyTask(ctx *iris.Context) {
	resp := types.RespPutTask{}

	uuid := ctx.GetString("uuid")
	if len(uuid) == 0 {
		resp.Errmsg = "'uuid' not set in URL path"
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}

	task := types.Task{}
	err := ctx.ReadJSON(&task)
	if err != nil {
		log.Printf("parse request body to task error: %v\n", err)
		resp.Errmsg = "parse request body to task error: " + err.Error()
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}

	updated, err := svc.UpdateTask(task)
	if err != nil {
		resp.Errmsg = "update task error: " + err.Error()
		if err == svc.ErrTaskNotFound {
			ctx.JSON(iris.StatusBadRequest, resp)
			return
		}
		ctx.JSON(iris.StatusInternalServerError, resp)
		return
	}

	resp.Success = true
	resp.Task = *updated
	ctx.JSON(iris.StatusOK, resp)
	return
}

// StartTask handles PUT /tasks/:uuid/start
// ==== Request ====
//
// ==== Response ====
func StartTask(ctx *iris.Context) {
	resp := types.RespStartTask{}

	uuid := ctx.GetString("uuid")
	if len(uuid) == 0 {
		resp.Errmsg = "'uuid' not set in URL path"
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}

	err := svc.StartTask(uuid)
	if err != nil {
		resp.Errmsg = "get task error: " + err.Error()
		if err == svc.ErrTaskNotFound {
			ctx.JSON(iris.StatusBadRequest, resp)
			return
		}
		ctx.JSON(iris.StatusInternalServerError, resp)
		return
	}

	resp.Success = true
	ctx.JSON(iris.StatusOK, resp)
	return
}

// StopTask handles PUT /tasks/:uuid/stop
// ==== Request ====
//
// ==== Response ====
func StopTask(ctx *iris.Context) {
	resp := types.RespStopTask{}

	uuid := ctx.GetString("uuid")

	if len(uuid) == 0 {
		resp.Errmsg = "'uuid' not set in URL path"
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
// ==== Request ====
// Body: Task
//
// ==== Response ====
// Body: Task
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
		resp.Errmsg = "'user_id' not set in request body"
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}
	if len(req.Name) == 0 {
		resp.Errmsg = "'name' not set in request body"
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}
	if len(req.LogDir) == 0 {
		resp.Errmsg = "'log_dir' not set in request body"
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}
	if req.CronInterval < 0 {
		resp.Errmsg = "'cron_interval' invalid in request body"
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
// ==== Request ====
//
// ==== Response ====
func DeleteTask(ctx *iris.Context) {
	resp := types.RespDeleteTask{}

	uuid := ctx.GetString("uuid")

	if len(uuid) == 0 {
		resp.Errmsg = "'uuid' not set in URL path"
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
