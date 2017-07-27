package main

import (
	"bitbucket.org/cdncache/taskmanager/api"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()

	app.Adapt(
		iris.DevLogger(),
		httprouter.New())

	app.Any("/", api.GetRoot)

	v1 := app.Party("/v1")

	v1.Get("/tasks", api.GetTasks)
	v1.Get("/tasks/:uuid", api.GetTask)
	v1.Post("/tasks", api.CreateTask)
	v1.Put("/tasks/:uuid", api.ModifyTask)
	v1.Put("/tasks/:uuid/start", api.StartTask)
	v1.Put("/tasks/:uuid/stop", api.StopTask)
	v1.Delete("/tasks/:uuid", api.DeleteTask)

	app.Listen(":8082")
}
