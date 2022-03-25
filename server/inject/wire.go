//go:build wireinject
// +build wireinject

package inject

import (
	"github.com/google/wire"
	"server/controller/backstage"
	"server/controller/front"
	"server/dao"
	"server/service"
)

//后台控制器
//------------------------------------------------------------------------------------------

//InitBackstageArticleController 初始化后台的文章控制器
func InitBackstageArticleController() *backstage.ArticleController {
	wire.Build(dao.NewArticleDAO, service.NewArticleService, backstage.NewArticleController)
	return nil
}

//InitBackstageCategoryController 初始化后台的分类控制器
func InitBackstageCategoryController() *backstage.CategoryController {
	wire.Build(dao.NewCategoryDAO, service.NewCategoryService, backstage.NewCategoryController)
	return nil
}

//InitBackstageTagController 初始化后台的标签控制器
func InitBackstageTagController() *backstage.TagController {
	wire.Build(dao.NewTagDAO, service.NewTagService, backstage.NewTagController)
	return nil
}

//InitBackstageUserController 初始化后台的用户控制器
func InitBackstageUserController() *backstage.UserController {
	wire.Build(dao.NewUserDAO, service.NewUserService, backstage.NewUserController)
	return nil
}

//前台控制器
//------------------------------------------------------------------------------------------

//InitFrontArticleController 初始化前台的文章控制器
func InitFrontArticleController() *front.ArticleController {
	wire.Build(dao.NewArticleDAO, service.NewArticleService, front.NewArticleController)
	return nil
}

//InitFrontArticleController 初始化前台的分类控制器
func InitFrontCategoryController() *front.CategoryController {
	wire.Build(dao.NewCategoryDAO, service.NewCategoryService, front.NewCategoryController)
	return nil
}

//InitFrontTagController 初始化前台的标签控制器
func InitFrontTagController() *front.TagController {
	wire.Build(dao.NewTagDAO, service.NewTagService, front.NewTagController)
	return nil
}

//InitFrontUserController 初始化前台的用户控制器
func InitFrontUserController() *front.UserController {
	wire.Build(dao.NewUserDAO, service.NewUserService, front.NewUserController)
	return nil
}
