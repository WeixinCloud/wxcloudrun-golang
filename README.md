# wxcloudrun-golang
[![GitHub license](https://img.shields.io/github/license/WeixinCloud/wxcloudrun-express)](https://github.com/WeixinCloud/wxcloudrun-express)
![GitHub package.json dependency version (prod)](https://img.shields.io/badge/golang-1.17.1-green)

## 简介
微信云托管 golang 模版，通过示例创建一张 todo_list 表，并对其进行增删改查的操作，对应 POST/DELETE/PUT/GET 四种请求的实现。

![](https://qcloudimg.tencent-cloud.cn/raw/3159427f92e66f3bd431c21e25f18793.png)


## 详细介绍：
1. 一键部署时将默认开通微信云托管中的 MySQL，并自动将数据库基本信息传入了环境变量中，可直接使用。（数据库信息获取及配置详情见：[init.go](https://github.com/WeixinCloud/wxcloudrun-golang/blob/main/db/init.go) ）。
2. [container.config.json](https://github.com/WeixinCloud/wxcloudrun-golang/blob/main/container.config.json) 仅用于在微信云托管中创建流水线时配套使用。
   * 如果不使用流水线，而是用本项目的代码在微信云托管控制台手动「新建版本」，则 container.config.json 配置文件不生效。最终版本部署效果以「新建版本」窗口中手动填写的值为准。
   * dataBaseName 和 executeSQLs 两个字段只有在服务第一次部署时生效，后续流水线触发的版本更新不会执行（避免重复初始化数据库）。
   

## 快速开始
前往 [微信云托管快速开始页面](https://developers.weixin.qq.com/miniprogram/dev/wxcloudrun/src/basic/guide.html)，选择相应语言的模板，根据引导完成部署。


## 目录结构
~~~
.
├── Dockerfile                Dockerfile 文件
├── LICENSE                   LICENSE 文件
├── README.md                 README 文件
├── container.config.json     微信云托管流水线配置
├── db                        数据库逻辑目录
├── go.mod                    go.mod 文件
├── go.sum                    go.sum 文件
├── index.html                主页 html 
├── main.go                   主函数入口
└── service                   接口服务逻辑目录
~~~

## 示例API列表

1 查询所有 todo 项

* URL路径：
  ```/api/todos```
  
* 请求示例：
```
curl -X GET  http://{ip}:{port}/api/todos
```

* 响应示例：
```
{
  "code": 0,
  "errorMsg": "",
  "data": [{
    "id": 1,
    "title": "工作1",
    "status": "准备中",
    "create_time": "2021-11-09T08:45:40Z",
    "update_time": "2021-11-09T08:45:40Z"
  }, {
    "id": 2,
    "title": "工作2",
    "status": "已开始",
    "create_time": "2021-11-09T08:46:11Z",
    "update_time": "2021-11-09T08:46:11Z"
  }]
}
```


2 根据ID查询 todo 项

* URL路径：
  ```/api/todos/:id```
  
* 请求示例：
```
curl -X GET  http://{ip}:{port}/api/todos/1
```

* 响应示例：
```
{
  "code": 0,
  "errorMsg": "",
  "data": {
    "id": 1,
    "title": "工作1",
    "status": "准备中",
    "create_time": "2021-11-09T08:45:40Z",
    "update_time": "2021-11-09T08:45:40Z"
  }
}
```


3 新增 todo 项目

* URL路径：
  ```/api/todos```
  
* 请求示例：
```
curl http://{ip}:{port}/api/todos \
  -X POST \
  -H 'Content-Type: application/json' \
  -d '{  
    "title":"工作1",
    "status":"准备中"
  }'
```

* 响应示例：
```
{
  "code": 0,
  "errorMsg": "",
  "data": {
    "id": 1,
    "title": "工作1",
    "status": "准备中",
    "create_time": "2021-11-09T08:45:40Z",
    "update_time": "2021-11-09T08:45:40Z"
  }
}
```

4 根据ID修改 todo 项目

* URL路径：
  ```/api/todos```
  
* 请求示例：
```
curl http://{ip}:{port}/api/todos \
  -X PUT \
  -H 'Content-Type: application/json' \
  -d '{  
    "id":1,
    "status":"已完成"
  }'
```

* 响应示例：
```
{
  "code": 0,
  "errorMsg": ""
}
```

5 根据ID删除 todo 项

* URL路径：
  ```/api/todos/:id```
  
* 请求示例：
```
curl http://{ip}:{port}/api/todos/1 \
  -X DELETE \
  -H 'Content-Type: application/json' \
  -d '{   }'
```

* 响应示例：
```
{
  "code": 0,
  "errorMsg": ""
}
```



## 简介
了解在微信云托管上如何用GO语言创建简单的http服务。通过示例创建一张todo_list表，并对其进行增删改查的操作，对应POST/DELETE/PUT/GET四种请求的实现。

版本：
Golang 1.17.1（修改版本，需要同步修改[Dockerfile](https://github.com/WeixinCloud/wxcloudrun-golang/blob/main/Dockerfile)中的基础镜像）

详细介绍：
1. 一键部署时将默认开通微信云托管中的MySQL，并自动将数据库基本信息传入了环境变量中，可直接使用。（数据库信息获取及配置详情见:[init.go](https://github.com/WeixinCloud/wxcloudrun-golang/blob/main/db/init.go)。）
2. [container.config.json](https://github.com/WeixinCloud/wxcloudrun-golang/blob/main/container.config.json)仅用于在微信云托管中创建流水线时配套使用。
   * 如果不使用流水线，而是用本项目的代码在微信云托管控制台手动「新建版本」，则container.config.json配置文件不生效。最终版本部署效果以「新建版本」窗口中手动填写的值为准。
   * 'dataBaseName'和‘executeSQLs’ 两个字段只有在服务第一次部署时生效，后续流水线触发的版本更新不会执行（避免重复初始化数据库）。
   
   
示例API列表：

1 查询所有 todo 项

* URL路径：
  ```/api/todos```
  
* 请求示例：
```
curl -X GET  http://{ip}:{port}/api/todos
```

* 响应示例：
```
{
	"code": 0,
	"errorMsg": "",
	"data": [{
		"id": 1,
		"title": "工作1",
		"status": "准备中",
		"create_time": "2021-11-09T08:45:40Z",
		"update_time": "2021-11-09T08:45:40Z"
	}, {
		"id": 2,
		"title": "工作2",
		"status": "已开始",
		"create_time": "2021-11-09T08:46:11Z",
		"update_time": "2021-11-09T08:46:11Z"
	}]
}
```


2 根据 ID 查询 todo 项

* URL路径：
  ```/api/todos/:id```
  
* 请求示例：
```
curl -X GET  http://{ip}:{port}/api/todos/1
```

* 响应示例：
```
{
	"code": 0,
	"errorMsg": "",
	"data": {
		"id": 1,
		"title": "工作1",
		"status": "准备中",
		"create_time": "2021-11-09T08:45:40Z",
		"update_time": "2021-11-09T08:45:40Z"
	}
}
```


3 新增 todo 项目

* URL路径：
  ```/api/todos```
  
* 请求示例：
```
curl http://{ip}:{port}/api/todos \
  -X POST \
  -H 'Content-Type: application/json' \
  -d '{  
      "title":"工作1",
      "status":"准备中"
  }'
```

* 响应示例：
```
{
	"code": 0,
	"errorMsg": "",
	"data": {
		"id": 1,
		"title": "工作1",
		"status": "准备中",
		"create_time": "2021-11-09T08:45:40Z",
		"update_time": "2021-11-09T08:45:40Z"
	}
}
```

4 根据 ID 修改 todo 项目

* URL路径：
  ```/api/todos```
  
* 请求示例：
```
curl http://{ip}:{port}/api/todos \
  -X PUT \
  -H 'Content-Type: application/json' \
  -d '{  
      "id":1,
      "status":"已完成"
  }'
```

* 响应示例：
```
{
	"code": 0,
	"errorMsg": ""
}
```

5 根据 ID 删除 todo 项

* URL路径：
  ```/api/todos/:id```
  
* 请求示例：
```
curl http://{ip}:{port}/api/todos/1 \
  -X DELETE \
  -H 'Content-Type: application/json' \
  -d '{   }'
```

* 响应示例：
```
{
	"code": 0,
	"errorMsg": ""
}
```

