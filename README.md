RocketMQ Producer Http Proxy

开发初衷：因为前端为PHP，而PHP官方暂无SDK, 为了考虑自身架构中不同语言的兼容性和开发成本，所以开发了这个小组件。
目前的版本是： beta 版本，还有很多可以做的，也许会存在某些问题（例如：消息丢失，消息回执，消息日志，性能等等），目前仅是内部使用，欢迎各位继续完善。

本项目是基于HTTP协议的 RocketMQ 前端生产者代理，目前仅仅提供了消息转发功能，可供其他语言端通过HTTP协议使用JSON格式快速便捷的生产消息。
至于消费，暂时没有HTTP协议的实现，以后有时间可以考虑基于HTTP2 的消息推送或者Pull模式的消息拉取。

Language: Golang ,version: v1.11.1

依赖客户端：

RocketMQ Client: https://github.com/apache/rocketmq-client-go

Fasthttp: http://github.com/valyala/fasthttp

Fasthttp Router: http://github.com/qiangxue/fasthttp-routing

