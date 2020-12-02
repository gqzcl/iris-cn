使用国内源“
https://developer.aliyun.com/mirror/NPM?from=tnpm

$ sudo npm install -g cnpm --registry=https://registry.npm.taobao.org
 
安装 
$ cnpm install [name]

同步模块
直接通过 sync 命令马上同步一个模块, 只有 cnpm 命令行才有此功能:

$ cnpm sync connect
当然, 你可以直接通过 web 方式来同步: /sync/connect

$ open https://npm.taobao.org/sync/connect