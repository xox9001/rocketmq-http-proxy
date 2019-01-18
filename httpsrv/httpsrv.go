package httpsrv

import (
	"config"
	"encoding/json"
	"github.com/qiangxue/fasthttp-routing"
)

var (
	AcquireMessageApi config.Api
)

func init(){
	AcquireMessageApi.Path = "/acquire_msg"
	AcquireMessageApi.Method = "post"
	AcquireMessageApi.Handle = func(ctx *routing.Context) error {

		var reqBody = ctx.Request.Body()
		var responseData = new(config.ResponseDataFormat)
		var rep []byte
		//默认为失败
		responseData.Msg = "[Error]:Data Is Empty,Please Check."

		//检查是否有数据
		if len(reqBody) > 0 {

			//默认应答为正常
			responseData.Status = 1
			responseData.Msg = "Acquire Message Success"
			HttpMssage := new(config.HttpMessage)
			err := json.Unmarshal(reqBody,HttpMssage)

			if err != nil {
				responseData.Status = 0
				responseData.Msg = "[Error]:Data Format Error,Please Check."
				goto RESPON
			}

			config.HttpMssageChannel <- HttpMssage.Decode()
		}

		goto RESPON

RESPON:
	rep,_ = json.Marshal(responseData)
	//TODO: 公共配置
	ctx.Response.Header.Set("Content-Type","application/json")
	ctx.Write(rep)

		return nil
	}
}