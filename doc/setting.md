环境：
5.8.18-1-MANJARO

# 配置 go

## 安装:

```bash
$ sudo pacman -S go
$ sudo pacman -S vscode
```
## 将go放入path路径：

```
$ sudo nano ~/zshrc

PATH="$HOME/go/bin:$PATH"

$ source ~/zshrc
```
##  在vscode配置go

## 配置环境

### 第一步：

参考 https://github.com/goproxy/goproxy.cn/blob/master/README.zh-CN.md

- 方式一

```
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
```

- 方式二 
```
$ export GO111MODULE=on
$ export GOPROXY=https://goproxy.cn
```

### 第二步：

先重启一遍vscode
然后按下Ctrl+Shift+P，输入 go:install,全部选中，安装
安装完成后就大功告成啦

### 第三步：配置IRIS

参考 https://github.com/kataras/iris/wiki/Installation

新建文件 go.mod

```go
module your_project_name

go 1.14

require (
    github.com/kataras/iris/v12 v12.1.8
)
```

然后：

```
$ go build
```

## 最后创建一个IRIS服务吧：

- 新建 main.go

```go
package main

import "github.com/kataras/iris/v12"

func main() {
    app := iris.New()
    app.RegisterView(iris.HTML("./views", ".html"))

    app.Get("/", func(ctx iris.Context) {
        ctx.ViewData("message", "Hello world!")
        ctx.View("hello.html")
    })

    app.Get("/user/{id:uint64}", func(ctx iris.Context) {
        userID, _ := ctx.Params().GetUint64("id")
        ctx.Writef("User ID: %d", userID)
    })

    app.Listen(":8080")
}
```

- 新建文件夹 views ，在 views 中新建 hello.html:

```HTML
<!-- file: ./views/hello.html -->
<html>
<head>
    <title>Hello Page</title>
</head>
<body>
    <h1>{{.message}}</h1>
</body>
</html>
```

- 最后：

```
$ go run main.go
```

在网页中打开 http://localhost:8080

就能看到一个 Hello World! 了。

