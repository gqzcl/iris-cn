package app

import (
	"iris-cn/controllers/api"

	"github.com/go-resty/resty/v2"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"github.com/sirupsen/logrus"
)

// InitIris 是初始化函数
func InitIris() {
	app := iris.New()
	app.Logger().SetLevel("warn")
	//恢复
	app.Use(recover.New())
	//
	app.Use(logger.New())
	//跨域中间件
	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		MaxAge:           600,
		AllowedMethods:   []string{iris.MethodGet, iris.MethodPost, iris.MethodOptions, iris.MethodHead, iris.MethodDelete, iris.MethodPut},
		AllowedHeaders:   []string{"*"},
	}))
	// handler := cors.Default().Handler(app)
	// handler = c.Handler(handler)
	app.AllowMethods(iris.MethodOptions)

	//api
	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Party("/article").Handle(new(api.Article))
		m.Party("/login").Handle(new(api.Login))
		m.Party("/comment").Handle(new(api.Comment))
		m.Party("/favorite").Handle(new(api.Favorite))
		m.Party("/tag").Handle(new(api.Tag))
		m.Party("/user").Handle(new(api.User))
	})

	// Simple HTTP and REST client library for Go
	app.Get("/api/img/proxy", func(i iris.Context) {
		url := i.FormValue("url")
		resp, err := resty.New().R().Get(url)
		i.Header("Content-Type", "image/jpg")
		if err == nil {
			_, _ = i.Write(resp.Body())
		} else {
			logrus.Error(err)
		}
	})
	app.Listen(":8080")
}
