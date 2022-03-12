//go:build wireinject
// +build wireinject

package inject

import (
	"github.com/google/wire"
	"server/controller/backstage"
	"server/controller/common"
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

//公共控制器
//------------------------------------------------------------------------------------------

//InitCommonArticleController 初始化公共的文章控制器
func InitCommonArticleController() *common.ArticleController {
	wire.Build(dao.NewArticleDAO, service.NewArticleService, common.NewArticleController)
	return nil
}

//InitCommonArticleController 初始化公共的分类控制器
func InitCommonCategoryController() *common.CategoryController {
	wire.Build(dao.NewCategoryDAO, service.NewCategoryService, common.NewCategoryController)
	return nil
}

//InitCommonTagController 初始化公共的标签控制器
func InitCommonTagController() *common.TagController {
	wire.Build(dao.NewTagDAO, service.NewTagService, common.NewTagController)
	return nil
}

//InitCommonUserController 初始化公共的用户控制器
func InitCommonUserController() *common.UserController {
	wire.Build(dao.NewUserDAO, service.NewUserService, common.NewUserController)
	return nil
}
