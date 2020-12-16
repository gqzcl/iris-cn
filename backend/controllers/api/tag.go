package api

import (
	"iris-cn/cache"
	"iris-cn/controllers/render"
	"iris-cn/model/constants"
	"iris-cn/services"

	"github.com/gqzcl/simple"
	"github.com/kataras/iris/v12"
)

type Tag struct {
	Ctx iris.Context
}

// 标签详情
func (c *Tag) GetBy(tagId int64) *simple.JsonResult {
	tag := cache.TagCache.Get(tagId)
	if tag == nil {
		return simple.JsonErrorMsg("标签不存在")
	}
	return simple.JsonData(render.BuildTag(tag))
}

// 标签列表
func (c *Tag) GetTags() *simple.JsonResult {
	page := simple.FormValueIntDefault(c.Ctx, "page", 1)
	tags, paging := services.TagService.FindPageByCnd(simple.NewSqlCnd().
		Eq("status", constants.StatusOk).
		Page(page, 200).Desc("id"))

	return simple.JsonPageData(render.BuildTags(tags), paging)
}

// 标签自动完成
func (c *Tag) PostAutocomplete() *simple.JsonResult {
	input := c.Ctx.FormValue("input")
	tags := services.TagService.Autocomplete(input)
	return simple.JsonData(tags)
}
