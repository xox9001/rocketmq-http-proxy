package config

import (
	"fmt"
	"github.com/apache/rocketmq-client-go/core"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"strings"
)


type ProducerConfig struct {
	NamesrvAddr 	string
	GroupID 		string
	InstanceName 	string
	GroupName       string
}

//服务启动函数，这里可以增加一个致命信息通道，用于在出错时友好的退出所有的服务
//TODO: 增加Result Channel 用户记录消息转发日志，暂时不实现
//func (p *ProducerConfig)Start(MsgChan <- chan *rocketmq.Message,Result chan <- *rocketmq.SendResult)(int){
func (p *ProducerConfig)Start(MsgChan <- chan *rocketmq.Message)(int){
	//TODO: check params
	config := new(rocketmq.ProducerConfig)
	config.GroupID = p.GroupID
	config.NameServer = p.NamesrvAddr
	config.InstanceName = p.InstanceName
	config.GroupName = p.GroupName

	producer,error := rocketmq.NewProducer(config)

	if error != nil {
		//TODO: Stop Srv
		panic(error)
	}

	producer.Start()
	defer producer.Shutdown()
	fmt.Printf("[Success]Producer: %s started... \n", producer)

	for msg := range MsgChan {
		fmt.Println(msg)
		_,err := producer.SendMessageSync(msg)
		//selector := queueSelectorByOrderID{}
		//result,err := producer.SendMessageOrderly(msgT,selector,1,3)

		if err != nil {
			fmt.Println(err)
		}else{
			fmt.Println("Msg Sync Send Success")
		}

		//Result <- result
	}

	return 0
}


type HttpSrvConfig struct {
	Addr string
	Concurrency int
	Srv  *fasthttp.Server
	Router *routing.Router
}

func (httpsrv *HttpSrvConfig)Init(){
	httpsrv.Router = routing.New()
	httpsrv.Srv = &fasthttp.Server{
		Concurrency: httpsrv.Concurrency,
		Handler:httpsrv.Router.HandleRequest,
	}
}

func (httpsrv *HttpSrvConfig)RegApi(api Api){
	method := strings.ToLower(api.Method)
	path := strings.ToLower(api.Path)

	switch method {
		case "get":
			httpsrv.Router.Get(path,api.Handle)
		case "post":
			httpsrv.Router.Post(path,api.Handle)
		default:
			httpsrv.Router.Any(path,api.Handle)
	}
	fmt.Println("API:",path," Register Success.")
}

func (httpsrv *HttpSrvConfig)Start(){
	fmt.Println("HTTP服务已启动,Listen:",httpsrv.Addr)
	panic(HttpSrv.Srv.ListenAndServe(httpsrv.Addr))
}


type ApiHandle func(ctx *routing.Context) error

type Api struct {
	Path string
	Method string
	Handle routing.Handler
}

//Json struct
type HttpMessage struct {
	Topic string    `json:"topic"`
	Body string		`json:"body"`
	DelayLevel int  `json:"delay_level"`
	Key string		`json:"key"`
	Tags string		`json:"tags"`
}

//TODO:这里暂时省略了信息验证，假设字段均合法的情况下
func (m *HttpMessage)Decode()*rocketmq.Message {
	rmqMsg := new(rocketmq.Message)
	rmqMsg.Topic = m.Topic
	rmqMsg.Body = m.Body
	rmqMsg.Keys = m.Key
	rmqMsg.Tags = m.Tags
	//messageDelayLevel=1s 5s 10s 30s 1m 2m 3m 4m 5m 6m 7m 8m 9m 10m 20m 30m 1h 2h
	if m.DelayLevel > 0 {
		rmqMsg.DelayTimeLevel = m.DelayLevel
	}

	return rmqMsg
}

type ResponseDataFormat struct {
	Status int `json:"status"`
	Msg string `json:msg`
	Data []interface{} `json:data,omitempty`
}