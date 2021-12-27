package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var writeData = func(w http.ResponseWriter, b []byte) (int, error) {
	return w.Write(b)
}

type Response struct {
	Message      string            `json:"message,omitempty"`
	Data         interface{}       `json:"data,omitempty"`
	StatusCode   int               `json:"-"`
	CustomHeader map[string]string `json:"-"`
}

//NewResponse init request
func NewResponse() Response {
	return Response{
		StatusCode:   http.StatusOK,
		Message:      http.StatusText(http.StatusOK),
		CustomHeader: map[string]string{},
	}
}

//SetError set error message
func (r Response) SetError(err error, code int) Response {
	r.StatusCode = code
	r.Message = err.Error()
	return r
}

//SetData set response data
func (r Response) SetData(data interface{}) Response {
	r.Data = data
	return r
}

//SetMessage init request
func (r Response) SetMessage(m string) Response {
	r.Message = m
	return r
}

func (r Response) SetCode(code int) Response {
	r.StatusCode = code
	return r
}

func (r Response) SetHeader(headerName string, value string) Response {
	r.CustomHeader[headerName] = value
	return r
}

func (r Response) send(w http.ResponseWriter) {
	b, err := json.Marshal(r)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "PUT,POST,GET,DELETE,PATCH,OPTION")

	for h, val := range r.CustomHeader {
		w.Header().Set(h, val)
	}

	w.WriteHeader(r.StatusCode)
	_, err = writeData(w, b)
	if err != nil {
		log.Println(err)
	}
}

//RespHandler response handler
func RespHandler(h Handle) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		resp := h(rw, r, p)
		resp.send(rw)
	}
}
