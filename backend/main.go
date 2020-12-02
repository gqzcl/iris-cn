package main

import (
	"flag"
	"iris-cn/app"
	"iris-cn/conf"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var configFile = flag.String("config", "./iriscn.yaml", "配置文件路径")

func init() {
	flag.Parse()
	//初始化配置
	conf := conf.Init(*configFile)

	//gorm
	gormConf := &gorm.Config{}

	//初始化日志
	if file, err := os.Open(conf.LogFile); err != nil {
		// if file, err := os.OpenFile(conf.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
		logrus.SetOutput(file)
		if conf.ShowSQL {
			gormConf.Logger = logger.New(log.New(file, "\r\n", log.LstdFlags), logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel:      logger.Info,
			})
		}
	} else {
		//logrus.Error(err)
		logrus.Error("打开日志失败")
	}

	//连接数据库
	if _, err := gorm.Open(mysql.Open(conf.MysqlURL), &gorm.Config{}); err != nil {
		logrus.Error(err)
	}
}
func main() {
	app.InitIris() // 实例一个iris对象
	// //配置路由
	// app.Get("/", func(ctx iris.Context) {
	// 	ctx.WriteString("Hello Iris")
	// })
	// app.Post("/", func(ctx iris.Context) {
	// 	ctx.Write([]byte("Hello Iris"))
	// })

	// // 路由分组
	// party := app.Party("/hello")
	// // 此处它的路由地址是: /hello/world
	// party.Get("/world", func(ctx iris.Context) {
	// 	ctx.WriteString("hello world")
	// })
	// userRouter := app.Party("/user")
	// // route: /user/{name}/home  例如:/user/dollarKiller/home
	// userRouter.Get("/{name:string}/home", func(ctx iris.Context) {
	// 	name := ctx.Params().Get("name")
	// 	ctx.Writef("you name: %s", name)
	// })
	// // route: /user/post
	// userRouter.Post("/post", func(ctx iris.Context) {
	// 	ctx.Writef("method:%s,path;%s", ctx.Method(), ctx.Path())
	// })

	// // 动态路由
	// // 路由传参
	// app.Get("/username/{name}", func(ctx iris.Context) {
	// 	name := ctx.Params().Get("name")
	// 	fmt.Println(name)
	// })

	// // 设置参数
	// app.Get("/profile/{id:int min(1)}", func(ctx iris.Context) {
	// 	i, e := ctx.Params().GetInt("id")
	// 	if e != nil {
	// 		ctx.WriteString("error you input")
	// 	}

	// 	ctx.WriteString(strconv.Itoa(i))
	// })

	// // 设置错误码
	// app.Get("/profile/{id:int min(1)}/friends/{friendid:int max(8) else 504}", func(ctx iris.Context) {
	// 	i, _ := ctx.Params().GetInt("id")
	// 	getInt, _ := ctx.Params().GetInt("friendid")
	// 	ctx.Writef("Hello id:%d looking for friend id: ", i, getInt)
	// }) // 如果没有传递所有路由的macros，这将抛出504错误代码而不是404.

	// // 正则表达式
	// app.Get("/lowercase/{name:string regexp(^[a-z]+)}", func(ctx iris.Context) {
	// 	ctx.Writef("name should be only lowercase, otherwise this handler will never executed: %s", ctx.Params().Get("name"))
	// })

	// // 启动服务器
	// app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
	// // 监听地址:本服务器上任意id端口8080,设置字符集utf8
}
