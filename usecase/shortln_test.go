package usecase

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"test/amartha/config"
	"test/amartha/database"
	"test/amartha/database/mocks"
	"test/amartha/router"
	"test/amartha/usecase/model"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestUsecase_GetShorten(t *testing.T) {
	tempHelperRandString := helperRandString
	defer func() {
		helperRandString = tempHelperRandString
	}()

	var mockDB *mocks.DBInterface
	type fields struct {
		Cfg config.Config
		DB  database.DB
	}
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		ps httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   router.Response
		patch  func()
	}{
		{
			name:   "err parse body",
			fields: fields{},
			args: args{
				r: &http.Request{
					Body: ioutil.NopCloser(strings.NewReader(`{`)),
				},
			},
			want:  router.NewResponse().SetError(fmt.Errorf("unexpected EOF"), http.StatusBadRequest),
			patch: func() {},
		},
		{
			name:   "err no url",
			fields: fields{},
			args: args{
				r: &http.Request{
					Body: ioutil.NopCloser(strings.NewReader(`{}`)),
				},
			},
			want:  router.NewResponse().SetError(fmt.Errorf("url is not present"), http.StatusBadRequest),
			patch: func() {},
		},
		{
			name:   "err invalid code",
			fields: fields{},
			args: args{
				r: &http.Request{
					Body: ioutil.NopCloser(strings.NewReader(`{
						"url":"some url",
						"shortcode":"code"
					}`)),
				},
			},
			want:  router.NewResponse().SetError(fmt.Errorf("The shortcode fails to meet requirement"), http.StatusUnprocessableEntity),
			patch: func() {},
		},
		{
			name:   "err exist shortcode",
			fields: fields{},
			args: args{
				r: &http.Request{
					Body: ioutil.NopCloser(strings.NewReader(`{
						"url":"some url",
						"shortcode":"codess"
					}`)),
				},
			},
			want: router.NewResponse().SetError(fmt.Errorf("The the desired shortcode is already in use"), http.StatusConflict),
			patch: func() {
				mockDB = new(mocks.DBInterface)
				mockDB.On("GetShortenByCode", "codess").Return(&model.ShortlnRequest{}).Once()
			},
		},
		{
			name:   "err save new shortcode",
			fields: fields{},
			args: args{
				r: &http.Request{
					Body: ioutil.NopCloser(strings.NewReader(`{
						"url":"some url",
						"shortcode":""
					}`)),
				},
			},
			want: router.NewResponse().SetError(fmt.Errorf("Some Error"), http.StatusInternalServerError),
			patch: func() {
				helperRandString = func(length int) string {
					return "codess"
				}

				mockDB = new(mocks.DBInterface)
				mockDB.On("GetShortenByCode", "codess").Return(&model.ShortlnRequest{}).Once()
				mockDB.On("GetShortenByCode", "codess").Return(nil).Once()
				mockDB.On("CreateShortenCode", &model.ShortlnRequest{
					Shortcode: "codess",
					Url:       "some url",
				}).Return(fmt.Errorf("Some Error")).Once()
			},
		},
		{
			name:   "Success save new shortcode",
			fields: fields{},
			args: args{
				r: &http.Request{
					Body: ioutil.NopCloser(strings.NewReader(`{
						"url":"some url",
						"shortcode":""
					}`)),
				},
			},
			want: router.NewResponse().SetData(&model.ShortlnResponse{Shortcode: "codess"}).SetCode(http.StatusCreated),
			patch: func() {
				helperRandString = func(length int) string {
					return "codess"
				}

				mockDB = new(mocks.DBInterface)
				mockDB.On("GetShortenByCode", "codess").Return(&model.ShortlnRequest{}).Once()
				mockDB.On("GetShortenByCode", "codess").Return(nil).Once()
				mockDB.On("CreateShortenCode", &model.ShortlnRequest{
					Shortcode: "codess",
					Url:       "some url",
				}).Return(nil).Once()
			},
		},
	}
	for _, tt := range tests {
		tt.patch()
		t.Run(tt.name, func(t *testing.T) {
			u := &Usecase{
				Cfg: tt.fields.Cfg,
				DB:  mockDB,
			}
			if got := u.GetShorten(tt.args.w, tt.args.r, tt.args.ps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.GetShorten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_GetURL(t *testing.T) {
	var mockDB *mocks.DBInterface
	type fields struct {
		Cfg config.Config
		DB  database.DBInterface
	}
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		ps httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   router.Response
		patch  func()
	}{
		{
			name:   "Not Found Code",
			fields: fields{},
			args: args{
				ps: httprouter.Params{
					httprouter.Param{
						Key:   "shorten",
						Value: "codess",
					},
				},
			},
			want: router.NewResponse().SetError(fmt.Errorf("The shortcode cannot be found in the system"), http.StatusNotFound),
			patch: func() {
				mockDB = new(mocks.DBInterface)
				mockDB.On("GetShortenByCode", "codess").Return(nil).Once()

			},
		},
		{
			name:   "Error Update Count",
			fields: fields{},
			args: args{
				ps: httprouter.Params{
					httprouter.Param{
						Key:   "shorten",
						Value: "codess",
					},
				},
			},
			want: router.NewResponse().SetError(fmt.Errorf("Some Error"), http.StatusInternalServerError),
			patch: func() {
				mockDB = new(mocks.DBInterface)
				mockDB.On("GetShortenByCode", "codess").Return(&model.ShortlnRequest{}).Once()
				mockDB.On("CountVisitingURL", "codess").Return(fmt.Errorf("Some Error")).Once()

			},
		},
		{
			name:   "Success Update Count",
			fields: fields{},
			args: args{
				ps: httprouter.Params{
					httprouter.Param{
						Key:   "shorten",
						Value: "codess",
					},
				},
			},
			want: router.NewResponse().SetCode(http.StatusFound).SetHeader("Location", ""),
			patch: func() {
				mockDB = new(mocks.DBInterface)
				mockDB.On("GetShortenByCode", "codess").Return(&model.ShortlnRequest{}).Once()
				mockDB.On("CountVisitingURL", "codess").Return(nil).Once()

			},
		},
	}
	for _, tt := range tests {
		tt.patch()
		t.Run(tt.name, func(t *testing.T) {
			u := &Usecase{
				Cfg: tt.fields.Cfg,
				DB:  mockDB,
			}
			if got := u.GetURL(tt.args.w, tt.args.r, tt.args.ps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.GetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_GetURLStats(t *testing.T) {
	var mockDB *mocks.DBInterface
	type fields struct {
		Cfg config.Config
		DB  database.DBInterface
	}
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		ps httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   router.Response
		patch  func()
	}{
		{
			name:   "Not Found Code",
			fields: fields{},
			args: args{
				ps: httprouter.Params{
					httprouter.Param{
						Key:   "shorten",
						Value: "codess",
					},
				},
			},
			want: router.NewResponse().SetError(fmt.Errorf("The shortcode cannot be found in the system"), http.StatusNotFound),
			patch: func() {
				mockDB = new(mocks.DBInterface)
				mockDB.On("GetShortenByCode", "codess").Return(nil).Once()

			},
		},
		{
			name:   "Found Count",
			fields: fields{},
			args: args{
				ps: httprouter.Params{
					httprouter.Param{
						Key:   "shorten",
						Value: "codess",
					},
				},
			},
			want: router.NewResponse().SetData(&model.ShortlnRequest{}),
			patch: func() {
				mockDB = new(mocks.DBInterface)
				mockDB.On("GetShortenByCode", "codess").Return(&model.ShortlnRequest{}).Once()

			},
		},
	}
	for _, tt := range tests {
		tt.patch()
		t.Run(tt.name, func(t *testing.T) {
			u := &Usecase{
				Cfg: tt.fields.Cfg,
				DB:  mockDB,
			}
			if got := u.GetURLStats(tt.args.w, tt.args.r, tt.args.ps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.GetURLStats() = %v, want %v", got, tt.want)
			}
		})
	}
}
