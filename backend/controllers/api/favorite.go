package api

import (
	"iris-cn/services"

	"github.com/gqzcl/simple"
	"github.com/kataras/iris/v12"
)

type Favorite struct {
	Ctx iris.Context
}

// 是否收藏了
func (c *Favorite) GetFavorited() *simple.JsonResult {
	user := services.UserTokenService.GetCurrent(c.Ctx)
	entityType := simple.FormValue(c.Ctx, "entityType")
	entityId := simple.FormValueInt64Default(c.Ctx, "entityId", 0)
	if user == nil || len(entityType) == 0 || entityId <= 0 {
		return simple.NewEmptyRspBuilder().Put("favorited", false).JsonResult()
	} else {
		tmp := services.FavoriteService.GetBy(user.Id, entityType, entityId)
		return simple.NewEmptyRspBuilder().Put("favorited", tmp != nil).JsonResult()
	}
}

// 取消收藏
func (c *Favorite) GetDelete() *simple.JsonResult {
	user := services.UserTokenService.GetCurrent(c.Ctx)
	entityType := simple.FormValue(c.Ctx, "entityType")
	entityId := simple.FormValueInt64Default(c.Ctx, "entityId", 0)
	if user == nil {
		return simple.JsonError(simple.ErrorNotLogin)
	}
	tmp := services.FavoriteService.GetBy(user.Id, entityType, entityId)
	if tmp != nil {
		services.FavoriteService.Delete(tmp.Id)
	}
	return simple.JsonSuccess()
}
