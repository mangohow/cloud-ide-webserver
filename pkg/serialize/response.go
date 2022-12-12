package serialize

import (
	code2 "github.com/mangohow/cloud-ide-webserver/pkg/code"
	"net/http"
	"sync"
)

type resResult struct {
	Data    interface{} `json:"data"`
	Status  uint32      `json:"status"`
	Message string      `json:"message"`
}

type Response struct {
	HttpStatus int
	R          resResult
}

var pool = sync.Pool{
	New: func() interface{} {
		return &Response{}
	},
}

func NewResponse(status int, code uint32, data interface{}) *Response {
	response := pool.Get().(*Response)
	response.HttpStatus = status
	response.R.Status = code
	response.R.Message = code2.GetMessage(code)
	response.R.Data = data

	return response
}

func PutResponse(res *Response) {
	if res != nil {
		res.R.Data = nil
		pool.Put(res)
	}

}

func NewResponseOk(code uint32, data interface{}) *Response {
	return NewResponse(http.StatusOK, code, data)
}

func NewResponseOKND(code uint32) *Response {
	return NewResponse(http.StatusOK, code, nil)
}
