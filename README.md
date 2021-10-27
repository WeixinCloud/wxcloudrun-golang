# wxcloudrun-golang
微信云托管 Go语言HTTP服务端示例

简介：了解在微信云托管上如何用GO语言创建简单的http服务。通过示例创建一张user表，并对其进行增删改查的操作，对应POST/DELETE/PUT/GET四种请求的实现。

版本：
Golang 1.17.1
* 修改版本，需要同步修改[Dockerfile](https://github.com/WeixinCloud/wxcloudrun-golang/blob/main/Dockerfile)中的基础镜像。

详细介绍：
1. 一键部署时将默认开通微信云托管中的MySQL，并自动将数据库基本信息传入了环境变量中，可直接使用。（数据库信息获取及配置详情见:[init.go](https://github.com/WeixinCloud/wxcloudrun-golang/blob/main/db/init.go)。）
2. [container.config.json](https://github.com/WeixinCloud/wxcloudrun-golang/blob/main/container.config.json)仅用于在微信云托管中创建流水线时配套使用。
   * 如果不使用流水线，而是用本项目的代码在微信云托管控制台手动「新建版本」，则container.config.json配置文件不生效。最终版本部署效果以「新建版本」窗口中手动填写的值为准。
   * 'dataBaseName'和‘executeSQLs’ 两个字段只有在服务第一次部署时生效，后续流水线触发的版本更新不会执行（避免重复初始化数据库）。
