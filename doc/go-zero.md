安装 go-zero
$ go get -u github.com/tal-tech/go-zero

安装 goctl
$ go get -u github.com/tal-tech/go-zero/tools/goctl

$ goctl api new app
$ cd app
$ go run app.go -f etc/app-api.yaml

安装protoc-gen-go
$ go get -u github.com/golang/protobuf/protoc-gen-go

安装etcd

分布式键值存储系统，可用于服务注册发现

https://github.com/etcd-io/etcd/releases

etcd --version
etcdctl version

echo 'export ETCDCTL_API=3' >> /etc/profile ## 环境变量添加 ETCDCTL_API=3[root@localhost etcd]# source /etc/profile # 是profile中修改的文件生效[root@localhost etcd]# ./etcdctl get mykey # 可以直接使用./etcdctl get key 命令了mykey

安装redis

$ sudo pacman -S redis

