package respx

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type Msg struct {
	Msg string `json:"msg"`
}

func Resp(w http.ResponseWriter, resp interface{}, msg interface{}) {
	body := Body{
		Code: 1,
		Msg:  "ok",
		Data: resp,
	}
	if msg.(Msg).Msg != "" {
		body.Msg = msg.(Msg).Msg
	}

	httpx.OkJson(w, body)
}
