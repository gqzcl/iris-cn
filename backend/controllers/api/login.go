package api

import (
	"iris-cn/auth"
	"iris-cn/controllers/render"
	"iris-cn/model"
	"iris-cn/services"

	"github.com/dchest/captcha"
	"github.com/gqzcl/simple"
	"github.com/kataras/iris/v12"
)

type Login struct {
	Ctx iris.Context
}

// 注册
func (c *Login) PostSignup() *simple.JsonResult {
	var (
		captchaId   = c.Ctx.PostValueTrim("captchaId")
		captchaCode = c.Ctx.PostValueTrim("captchaCode")
		email       = c.Ctx.PostValueTrim("email")
		username    = c.Ctx.PostValueTrim("username")
		password    = c.Ctx.PostValueTrim("password")
		rePassword  = c.Ctx.PostValueTrim("rePassword")
		nickname    = c.Ctx.PostValueTrim("nickname")
		ref         = c.Ctx.FormValue("ref")
	)
	if !captcha.VerifyString(captchaId, captchaCode) {
		return simple.JsonError(auth.CaptchaError)
	}
	user, err := services.UserService.SignUp(username, email, nickname, password, rePassword)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return c.GenerateLoginResult(user, ref)
}

// 用户名密码登录
func (c *Login) PostSignin() *simple.JsonResult {
	var (
		captchaId   = c.Ctx.PostValueTrim("captchaId")
		captchaCode = c.Ctx.PostValueTrim("captchaCode")
		username    = c.Ctx.PostValueTrim("username")
		password    = c.Ctx.PostValueTrim("password")
		ref         = c.Ctx.FormValue("ref")
	)
	if !captcha.VerifyString(captchaId, captchaCode) {
		return simple.JsonError(auth.CaptchaError)
	}
	user, err := services.UserService.SignIn(username, password)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return c.GenerateLoginResult(user, ref)
}

// 退出登录
func (c *Login) GetSignout() *simple.JsonResult {
	err := services.UserTokenService.Signout(c.Ctx)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonSuccess()
}

// user: login user, ref: 登录来源地址，需要控制登录成功之后跳转到该地址
func (c *Login) GenerateLoginResult(user *model.User, ref string) *simple.JsonResult {
	token, err := services.UserTokenService.Generate(user.Id)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.NewEmptyRspBuilder().
		Put("token", token).
		Put("user", render.BuildUser(user)).
		Put("ref", ref).JsonResult()
}
