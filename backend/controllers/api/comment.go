package api

import (
	"iris-cn/controllers/render"
	"iris-cn/model"
	"iris-cn/services"
	"strconv"

	"github.com/gqzcl/simple"
	"github.com/kataras/iris/v12"
)

type Comment struct {
	Ctx iris.Context
}

func (c *Comment) GetList() *simple.JsonResult {
	var (
		err        error
		cursor     int64
		entityType string
		entityId   int64
	)
	cursor = simple.FormValueInt64Default(c.Ctx, "cursor", 0)

	if entityType, err = simple.FormValueRequired(c.Ctx, "entityType"); err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	if entityId, err = simple.FormValueInt64(c.Ctx, "entityId"); err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	comments, cursor := services.CommentService.GetComments(entityType, entityId, cursor)
	return simple.JsonCursorData(render.BuildComments(comments), strconv.FormatInt(cursor, 10))
}

func (c *Comment) PostCreate() *simple.JsonResult {
	user := services.UserTokenService.GetCurrent(c.Ctx)
	if err := services.UserService.CheckPostStatus(user); err != nil {
		return simple.JsonError(err)
	}

	form := &model.CreateCommentForm{}
	err := simple.ReadForm(c.Ctx, form)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	comment, err := services.CommentService.Publish(user.Id, form)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	return simple.JsonData(render.BuildComment(*comment))
}
