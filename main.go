package main

import (
	"config"
	"httpsrv"
)

func main(){
	go config.Producer.Start(config.HttpMssageChannel)
	config.HttpSrv.Init()
	config.HttpSrv.RegApi(httpsrv.AcquireMessageApi)
	config.HttpSrv.Start()
}
