package respx

import (
	"fmt"
	"reflect"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ReturnBody(resp any) *Body {
	//data := remove
	fmt.Println(resp)
	v := reflect.ValueOf(resp).Elem()
	msg := v.FieldByName("RespMsg")
	fmt.Println(msg.FieldByName("Msg"))
	return &Body{
		Code: 1,
		Msg:  msg.String(),
		Data: resp,
	}
}
