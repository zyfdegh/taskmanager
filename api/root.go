package api

import (
	"gopkg.in/kataras/iris.v6"
)

// GetRoot handles GET /
func GetRoot(ctx *iris.Context) {
	ctx.WriteString("cdncache taskmanager server")
	return
}
