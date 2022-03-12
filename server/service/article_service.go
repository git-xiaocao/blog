package service

import "server/dao"

type ArticleService struct {
	dao *dao.ArticleDAO
}

func NewArticleService(dao *dao.ArticleDAO) *ArticleService {
	return &ArticleService{dao}
}
