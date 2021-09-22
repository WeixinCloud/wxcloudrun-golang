# wxcloudrun-golang
微信云托管 Go语言Hello World示例

简介：了解在微信云托管上如何用GO语言创建简单的 Hello World服务.

详细介绍：
1. 本示例中，使用的是Golang 1.17.1，服务通过80端口对外。
    * 在代码中修改端口号之后，如果使用流水线部署版本，请确保container.config.json中的containerPort字段也同步修改；如果在微信云托管控制台手动「新建版本」，请确保“监听端口”字段与代码中端口号保持一致，否则会引发部署失败。
2. 代码仓库中的container.config.json文件仅用于在微信云托管中创建流水线。如果不使用流水线，而是用本项目的代码在微信云托管控制台手动「新建版本」，则container.config.json配置文件不生效。最终版本部署效果以「新建版本」窗口中手动填写的值为准。
