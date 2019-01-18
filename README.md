RocketMQ Producer Http Proxy
Language: Golang 
version: v1.11.1

开发初衷：因为前端为PHP，而PHP官方暂无SDK, 为了考虑自身架构中不同语言的兼容性和开发成本，所以用了半天时间开发了这个小组件。

目前的版本是： beta 版本，还有很多可以做的，也许会存在某些问题（例如：消息丢失，消息回执，消息日志，性能等等），

目前仅是内部使用，欢迎各位继续完善。

本项目是基于HTTP协议的 RocketMQ 前端生产者代理，目前仅仅提供了消息转发功能，可供其他语言端通过HTTP协议使用JSON格式快速便捷的生产消息。
至于消费，暂时没有HTTP协议的实现，以后有时间可以考虑基于HTTP2 的消息推送或者Pull模式的消息拉取。

使用之前请先修改 config文件夹中 config.go的配置信息

//rocketmq namesrvaddr 地址

Producer.NamesrvAddr = "192.168.31.152:9876"

//需要监听的HTTP服务地址和端口

HttpSrv.Addr = "192.168.31.152:7776"

//Http服务最大并发数，具体作用请千万 fasthttp 配置查看，主要是http服务能最大同时处理的并发连接数；

//fasthttp 是单线程模型，有效避免内存复制和回收问题，性能比较高

HttpSrv.Concurrency = 10000

//消息处理通道，当http server接受到消息后，经过Decode 方法将消息转发到此channel，再由内部producer发送至broker,producer到broker为长连接，而且只在服务启动时请求namesrv一次

HttpMssageChannel = make(chan *rocketmq.Message,10000)

==========
目前只开放了一个消息转发接口

接口名称：/acquire_msg

Url： http://[server addr]:[port]/acquire_msg

Method: POST

Request Body 为json格式

  Topic string    `json:"topic"`
  
	Body string		`json:"body"`
  
	DelayLevel int  `json:"delay_level"`
  
	Key string		`json:"key"`
  
	Tags string		`json:"tags"`
  

Example: {"topic":"abcaa","body":"aaaaa","key":"1","tags":"show msg","delay_level":1}



Req Example:

`curl -X POST 
  http://192.168.31.152:7776/acquire_msg 
  -d '{"topic":"abcaa","body":"aaaaa","key":"1"}'
`

Response:
`
{
    "status": 1,
    "Msg": "Acquire Message Success",
    "Data": null
}
`

依赖客户端：

RocketMQ Client: https://github.com/apache/rocketmq-client-go

Fasthttp: http://github.com/valyala/fasthttp

Fasthttp Router: http://github.com/qiangxue/fasthttp-routing

