package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"log"
	"server/controller"
	"server/middleware"
)

func main() {
	app := iris.Default()

	mvc.Configure(app.Party("/api/article"), articleMvc)

	err := app.Run(iris.Addr(":44444"))

	if err != nil {
		log.Fatalln(err)
		return
	}
}

func articleMvc(app *mvc.Application) {
	app.Router.Use(middleware.Token)
	app.Handle(new(controller.ArticleController))
}
