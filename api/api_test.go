package api

import (
	"test/amartha/config"
	"test/amartha/router"
	"test/amartha/usecase"
	"testing"

	routeMock "test/amartha/router/mocks"

	"github.com/stretchr/testify/mock"
)

func TestAPI_Run(t *testing.T) {
	tempHandlerNewRouter := handlerNewRouter
	defer func() {
		handlerNewRouter = tempHandlerNewRouter
	}()

	type fields struct {
		Cfg        *config.Config
		Interactor usecase.Interactor
	}
	tests := []struct {
		name   string
		fields fields
		patch  func()
	}{
		{
			name: "InitFunction_SuccessOnly",
			fields: fields{
				Cfg: &config.Config{
					Server: &config.Server{
						Port: "1000",
					},
				},
				Interactor: &usecase.Usecase{},
			},
			patch: func() {
				handlerNewRouter = func() router.RouterIO {
					mr := new(routeMock.RouterIO)
					mr.On("POST", "/shorten", mock.Anything).Once()
					mr.On("GET", "/:shorten", mock.Anything).Once()
					mr.On("GET", "/:shorten/stats", mock.Anything).Once()
					mr.On("Serve", "1000").Once()

					return mr
				}
			},
		},
	}
	for _, tt := range tests {
		tt.patch()
		t.Run(tt.name, func(t *testing.T) {
			a := &API{
				Cfg:        tt.fields.Cfg,
				Interactor: tt.fields.Interactor,
			}
			a.Run()
		})
	}
}
