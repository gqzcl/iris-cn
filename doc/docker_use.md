环境：
5.8.18-1-MANJARO

启动docker：

开机自动启动
$ sudo systemctl enable docker

启动docker
$ sudo systemctl start docker

重启doker
$ sudo systemctl restart docker

检查是否安装成功：
$ sudo docker run hello-world

如果安装成功应该会看到Hello from Docker。

镜像加速：
参考：https://www.runoob.com/docker/docker-mirror-acceleration.html
推荐使用阿里云镜像加速器：https://cr.console.aliyun.com/cn-hangzhou/instances/mirrors

如果没有docker文件夹则新建一个
$ sudo mkdir -p /etc/docker

写入：

$ sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://<id>.mirror.aliyuncs.com"]
}
EOF

保存后重启服务：
$ sudo systemctl daemon-reload
$ sudo systemctl restart docker

下载镜像：
$ sudo docker pull ubuntu

也可以自己下别的，具体去 docker hub 看

运行镜像：
$ sudo docker run -t -i ubuntu /bin/bash

查看镜像信息：
$ cat /proc/version

退出： exit

列出本地镜像：
$ sudo docker images

创建Docker用户组：
$ sudo usermod -aG docker <username>
$ restart

查看info 
$ docker info

