package router

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestNewResponse(t *testing.T) {
	tests := []struct {
		name string
		want Response
	}{
		{
			name: "SuccessOnly",
			want: Response{
				StatusCode:   http.StatusOK,
				Message:      http.StatusText(http.StatusOK),
				CustomHeader: map[string]string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_SetError(t *testing.T) {
	type fields struct {
		Message      string
		Data         interface{}
		StatusCode   int
		CustomHeader map[string]string
	}
	type args struct {
		err  error
		code int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Response
	}{
		{
			name: "SuccessOnly",
			want: Response{},
			args: args{
				err: errors.New(""),
			},
			fields: fields{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Response{
				Message:      tt.fields.Message,
				Data:         tt.fields.Data,
				StatusCode:   tt.fields.StatusCode,
				CustomHeader: tt.fields.CustomHeader,
			}
			if got := r.SetError(tt.args.err, tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.SetError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_SetData(t *testing.T) {
	type fields struct {
		Message      string
		Data         interface{}
		StatusCode   int
		CustomHeader map[string]string
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Response
	}{
		{
			name: "SuccessOnly",
			fields: fields{
				Data: "",
			},
			args: args{},
			want: Response{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Response{
				Message:      tt.fields.Message,
				Data:         tt.fields.Data,
				StatusCode:   tt.fields.StatusCode,
				CustomHeader: tt.fields.CustomHeader,
			}
			if got := r.SetData(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.SetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_SetMessage(t *testing.T) {
	type fields struct {
		Message      string
		Data         interface{}
		StatusCode   int
		CustomHeader map[string]string
	}
	type args struct {
		m string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Response
	}{
		{
			name:   "SuccessOnly",
			fields: fields{},
			args:   args{},
			want:   Response{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Response{
				Message:      tt.fields.Message,
				Data:         tt.fields.Data,
				StatusCode:   tt.fields.StatusCode,
				CustomHeader: tt.fields.CustomHeader,
			}
			if got := r.SetMessage(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.SetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_SetCode(t *testing.T) {
	type fields struct {
		Message      string
		Data         interface{}
		StatusCode   int
		CustomHeader map[string]string
	}
	type args struct {
		code int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Response
	}{
		{

			name:   "SuccessOnly",
			fields: fields{},
			args:   args{},
			want:   Response{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Response{
				Message:      tt.fields.Message,
				Data:         tt.fields.Data,
				StatusCode:   tt.fields.StatusCode,
				CustomHeader: tt.fields.CustomHeader,
			}
			if got := r.SetCode(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.SetCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_SetHeader(t *testing.T) {
	type fields struct {
		Message      string
		Data         interface{}
		StatusCode   int
		CustomHeader map[string]string
	}
	type args struct {
		headerName string
		value      string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Response
	}{
		{
			name: "SuccessOnly",
			fields: fields{
				CustomHeader: map[string]string{},
			},
			args: args{},
			want: Response{
				CustomHeader: map[string]string{
					"": "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Response{
				Message:      tt.fields.Message,
				Data:         tt.fields.Data,
				StatusCode:   tt.fields.StatusCode,
				CustomHeader: tt.fields.CustomHeader,
			}
			if got := r.SetHeader(tt.args.headerName, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.SetHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRespHandler(t *testing.T) {
	type args struct {
		h Handle
	}
	tests := []struct {
		name string
		args args
		want httprouter.Handle
	}{
		{
			name: "SuccessOnly",
			args: args{
				h: func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) Response {
					return Response{}
				},
			},
			want: func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RespHandler(tt.args.h); reflect.TypeOf(tt.want) != reflect.TypeOf(got) {
				t.Errorf("RespHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_send(t *testing.T) {
	tempWriteData := writeData
	defer func() {
		writeData = tempWriteData
	}()

	type fields struct {
		Message      string
		Data         interface{}
		StatusCode   int
		CustomHeader map[string]string
	}
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		patch  func()
	}{
		{
			name: "SuccessOnly",
			fields: fields{
				CustomHeader: map[string]string{
					"": "",
				},
			},
			args: args{
				w: &httptest.ResponseRecorder{},
			},
			patch: func() {
				writeData = func(w http.ResponseWriter, b []byte) (int, error) {
					return 0, nil
				}
			},
		},
		{
			name: "ErrorWrite",
			fields: fields{
				CustomHeader: map[string]string{
					"": "",
				},
			},
			args: args{
				w: &httptest.ResponseRecorder{},
			},
			patch: func() {
				writeData = func(w http.ResponseWriter, b []byte) (int, error) {
					return 0, fmt.Errorf("some err")
				}
			},
		},
	}
	for _, tt := range tests {
		tt.patch()
		t.Run(tt.name, func(t *testing.T) {
			r := Response{
				Message:      tt.fields.Message,
				Data:         tt.fields.Data,
				StatusCode:   tt.fields.StatusCode,
				CustomHeader: tt.fields.CustomHeader,
			}
			r.send(tt.args.w)
		})
	}
}

func Test_writeData(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Success Only",
			args: args{
				w: &httptest.ResponseRecorder{},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := writeData(tt.args.w, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("writeData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("writeData() = %v, want %v", got, tt.want)
			}
		})
	}
}
