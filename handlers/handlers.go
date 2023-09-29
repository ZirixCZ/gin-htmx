package handlers

import (
	"github.com/kataras/iris/v12"
)

func RegisterHandler(app *iris.Application) {
	app.Get("/", func(ctx iris.Context) {
		ctx.View("login.html")
	})
}
