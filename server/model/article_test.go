package model

import (
	"fmt"
	"testing"
)

func TestQueryArticles(t *testing.T) {

	err, articles := Article{}.QueryList(0, 10)
	if err != nil {
		return
	}

	fmt.Printf("%+#v\n", articles)
}

func TestAddArticle(t *testing.T) {
	article := Article{
		Title:    "标题呀",
		Content:  "内容呀",
		Category: Category{Id: 1},
		Tags: []*Tag{
			{Id: 4},
			{Id: 5},
			{Name: "new tag"},
		},
	}
	err := article.Add()
	if err != nil {
		panic(err)
	}

}

func TestUpdateArticle(t *testing.T) {
	article := Article{
		Title:    "标题呀",
		Content:  "内容呀",
		Category: Category{Id: 1},
		Tags: []*Tag{
			{Id: 4},
			{Id: 5},
		},
	}
	article.Id = 3
	err := article.Update()
	if err != nil {
		panic(err)
	}

}

func TestQueryArticlesByTagIds(t *testing.T) {

	err, articles := Article{}.QueryByTagIds([]int64{1, 2, 3, 4}, 0, 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+#v\n", articles)
}
