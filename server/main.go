package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"log"
	"server/inject"
	"server/middleware"
)

func main() {

	app := iris.Default()

	baseParty := app.Party("/api")

	InitBackstageMVC(baseParty.Party("/backstage"))

	InitCommonMVC(baseParty.Party("/common"))

	err := app.Run(iris.Addr(":44444"))

	if err != nil {
		log.Fatalln(err)
		return
	}
}

// InitBackstageMVC 初始化后台MVC
func InitBackstageMVC(party iris.Party) {
	party.Use(middleware.Token)
	mvc.Configure(party.Party("/article"), func(app *mvc.Application) {
		app.Handle(inject.InitBackstageArticleController())
	})
	mvc.Configure(party.Party("/category"), func(app *mvc.Application) {
		app.Handle(inject.InitBackstageCategoryController())
	})
	mvc.Configure(party.Party("/tag"), func(app *mvc.Application) {
		app.Handle(inject.InitBackstageTagController())
	})
	mvc.Configure(party.Party("/user"), func(app *mvc.Application) {
		app.Handle(inject.InitBackstageTagController())
	})
}

// InitCommonMVC 初始化公共MVC
func InitCommonMVC(party iris.Party) {
	party.Use(middleware.Token)
	mvc.Configure(party.Party("/article"), func(app *mvc.Application) {
		app.Handle(inject.InitCommonArticleController())
	})
	mvc.Configure(party.Party("/category"), func(app *mvc.Application) {
		app.Handle(inject.InitCommonCategoryController())
	})
	mvc.Configure(party.Party("/tag"), func(app *mvc.Application) {
		app.Handle(inject.InitCommonTagController())
	})
	mvc.Configure(party.Party("/user"), func(app *mvc.Application) {
		app.Handle(inject.InitBackstageTagController())
	})
}
