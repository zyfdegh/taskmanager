package midware

import (
	"log"

	"gopkg.in/kataras/iris.v6"
)

func ApiMidware(ctx *iris.Context) {
	log.Printf("Handling %s %s...", ctx.Method(), ctx.Path())
	ctx.Next()
}
