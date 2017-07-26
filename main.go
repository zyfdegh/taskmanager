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

	app.Get("/", api.GetRoot)

	app.Get("/tasks", api.GetTasks)
	app.Get("/tasks/:id", api.GetTask)
	app.Post("/tasks", api.CreateTask)
	app.Put("/tasks/:id", api.ModifyTask)
	app.Delete("/tasks/:id", api.DeleteTask)

	app.Listen(":8082")
}
