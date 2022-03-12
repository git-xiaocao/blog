// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package inject

import (
	"server/controller/backstage"
	"server/controller/common"
	"server/dao"
	"server/service"
)

// Injectors from wire.go:

//InitBackstageArticleController 初始化后台的文章控制器
func InitBackstageArticleController() *backstage.ArticleController {
	articleDAO := dao.NewArticleDAO()
	articleService := service.NewArticleService(articleDAO)
	articleController := backstage.NewArticleController(articleService)
	return articleController
}

//InitBackstageCategoryController 初始化后台的分类控制器
func InitBackstageCategoryController() *backstage.CategoryController {
	categoryDAO := dao.NewCategoryDAO()
	categoryService := service.NewCategoryService(categoryDAO)
	categoryController := backstage.NewCategoryController(categoryService)
	return categoryController
}

//InitBackstageTagController 初始化后台的标签控制器
func InitBackstageTagController() *backstage.TagController {
	tagDAO := dao.NewTagDAO()
	tagService := service.NewTagService(tagDAO)
	tagController := backstage.NewTagController(tagService)
	return tagController
}

//InitBackstageUserController 初始化后台的用户控制器
func InitBackstageUserController() *backstage.UserController {
	userDAO := dao.NewUserDAO()
	userService := service.NewUserService(userDAO)
	userController := backstage.NewUserController(userService)
	return userController
}

//InitCommonArticleController 初始化公共的文章控制器
func InitCommonArticleController() *common.ArticleController {
	articleDAO := dao.NewArticleDAO()
	articleService := service.NewArticleService(articleDAO)
	articleController := common.NewArticleController(articleService)
	return articleController
}

//InitCommonArticleController 初始化公共的分类控制器
func InitCommonCategoryController() *common.CategoryController {
	categoryDAO := dao.NewCategoryDAO()
	categoryService := service.NewCategoryService(categoryDAO)
	categoryController := common.NewCategoryController(categoryService)
	return categoryController
}

//InitCommonTagController 初始化公共的标签控制器
func InitCommonTagController() *common.TagController {
	tagDAO := dao.NewTagDAO()
	tagService := service.NewTagService(tagDAO)
	tagController := common.NewTagController(tagService)
	return tagController
}

//InitCommonUserController 初始化公共的用户控制器
func InitCommonUserController() *common.UserController {
	userDAO := dao.NewUserDAO()
	userService := service.NewUserService(userDAO)
	userController := common.NewUserController(userService)
	return userController
}
