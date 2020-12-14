package api

import (
	"iris-cn/service"

	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
)

type Article struct {
	Ctx iris.Context
}

func (c *Article) Getby(articleId int64) *simple.JsonResult {
	article := service.ArticleService.Get(articleId)
}
