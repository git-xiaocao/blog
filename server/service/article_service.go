package service

import (
	"errors"
	"server/dao"
	"server/dto"
	"server/model"
)

type ArticleService struct {
	dao *dao.ArticleDAO
}

func (a *ArticleService) Add(dto dto.ArticleAddDTO) (err error) {
	if dto.Title == "" {
		err = errors.New("字段title为空")
		return
	}
	if dto.Content == "" {
		err = errors.New("字段content为空")
		return
	}
	if dto.PathMark == "" {
		err = errors.New("字段pathMark为空")
		return
	}

	var tags []*model.Tag

	for _, item := range dto.TagIds {
		tags = append(tags, &model.Tag{Id: item})
	}

	for _, item := range dto.NewTags {
		tags = append(tags, &model.Tag{Name: item})
	}

	article := model.Article{
		CategoryId: dto.CategoryId,
		Title:      dto.Title,
		Content:    dto.Content,
		PathMark:   dto.PathMark,
		Tags:       tags,
	}
	err = a.dao.Add(article)

	return
}

func (a *ArticleService) Update(dto dto.ArticleUpdateDTO) (err error) {
	if dto.Title == "" {
		err = errors.New("字段title为空")
		return
	}
	if dto.Content == "" {
		err = errors.New("字段content为空")
		return
	}
	if dto.PathMark == "" {
		err = errors.New("字段pathMark为空")
		return
	}

	var tags []*model.Tag

	for _, item := range dto.TagIds {
		tags = append(tags, &model.Tag{Id: item})
	}

	for _, item := range dto.NewTags {
		tags = append(tags, &model.Tag{Name: item})
	}

	article := model.Article{
		CategoryId: dto.CategoryId,
		Title:      dto.Title,
		Content:    dto.Content,
		PathMark:   dto.PathMark,
		Tags:       tags,
	}
	article.Id = dto.Id
	err = a.dao.Update(article)
	return
}

func (a *ArticleService) Delete(dto dto.ArticleDeleteDTO) (err error) {
	err = a.dao.DeleteById(dto.Id)
	return
}

func (a *ArticleService) Get(pathMark string) (result model.Article, err error) {
	result, err = a.dao.QueryByPathMark(pathMark)
	return
}

func (a *ArticleService) Exist(pathMark string) (result bool) {
	result = a.dao.ExistByPathMark(pathMark)
	return
}

func (a *ArticleService) List(tagIds []int64, offset, limit int) (result []*model.Article, err error) {
	if len(tagIds) == 0 {
		result, err = a.dao.QueryList(offset, limit)
	} else {
		result, err = a.dao.QueryListByTagIds(tagIds, offset, limit)
	}
	return
}

func NewArticleService(dao *dao.ArticleDAO) *ArticleService {
	return &ArticleService{dao}
}
