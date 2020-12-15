package api

import (
	"iris-cn/model/constants"
	"iris-cn/services"

	"github.com/gqzcl/simple"
	"github.com/kataras/iris/v12"
)

type Article struct {
	Ctx iris.Context
}

//文章详情
func (c *Article) Getby(articleId int64) *simple.JsonResult {
	article := service.ArticleService.Get(articleId)
	if article == nil || article.Status == constants.StatusDeleted {
		return simple.JsonErrorCode(404, "文章不存在")
	}

	user := services.UserTokenService.GetCurrent(c.Ctx)
	if user != nil {
		if article.UserId != user.Id && article.Status == constants.StatusPending {
			return simple.JsonErrorCode(403, "文章审核中")
		}
	} else {
		if article.Status == constants.StatusPending {
			return simple.JsonErrorCode(403, "文章审核中")
		}
	}
	// 增加浏览量
	services.ArticleService.IncrViewCount(articleId)

	return simple.JsonData(render.BuildArticle(article))
}
