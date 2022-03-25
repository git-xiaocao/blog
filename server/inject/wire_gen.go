// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package inject

import (
	"server/controller/backstage"
	"server/controller/front"
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

//InitFrontArticleController 初始化前台的文章控制器
func InitFrontArticleController() *front.ArticleController {
	articleDAO := dao.NewArticleDAO()
	articleService := service.NewArticleService(articleDAO)
	articleController := front.NewArticleController(articleService)
	return articleController
}

//InitFrontArticleController 初始化前台的分类控制器
func InitFrontCategoryController() *front.CategoryController {
	categoryDAO := dao.NewCategoryDAO()
	categoryService := service.NewCategoryService(categoryDAO)
	categoryController := front.NewCategoryController(categoryService)
	return categoryController
}

//InitFrontTagController 初始化前台的标签控制器
func InitFrontTagController() *front.TagController {
	tagDAO := dao.NewTagDAO()
	tagService := service.NewTagService(tagDAO)
	tagController := front.NewTagController(tagService)
	return tagController
}

//InitFrontUserController 初始化前台的用户控制器
func InitFrontUserController() *front.UserController {
	userDAO := dao.NewUserDAO()
	userService := service.NewUserService(userDAO)
	userController := front.NewUserController(userService)
	return userController
}
