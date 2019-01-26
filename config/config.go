package config

import "github.com/apache/rocketmq-client-go/core"

var (
	Producer ProducerConfig
	HttpSrv HttpSrvConfig
	HttpMssageChannel chan *rocketmq.Message
)

func init(){
	//TODO: 可优化至配置文件中读取，做个标记后续完善
	Producer.NamesrvAddr = "10.1.1.131:9876"
	Producer.GroupID = "http-proxy"
	Producer.InstanceName = "version:1.0"
	Producer.GroupName = "http-proxy-producer"
	HttpSrv.Addr = "10.1.1.131:7776"
	HttpSrv.Concurrency = 10000
	HttpMssageChannel = make(chan *rocketmq.Message,10000)
}