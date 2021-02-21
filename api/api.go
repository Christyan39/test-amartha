package api

import (
	"log"
	"test/amartha/config"
	"test/amartha/router"
	"test/amartha/usecase"
)

type API struct {
	Cfg        *config.Config
	Interactor usecase.Interactor
}

var handlerNewRouter = router.NewRouter

func (a *API) Run() {
	log.Println("==============Available Endpoints===============")
	router := handlerNewRouter()
	router.POST("/shorten", a.Interactor.GetShorten)
	router.GET("/:shorten", a.Interactor.GetURL)
	router.GET("/:shorten/stats", a.Interactor.GetURLStats)

	log.Println("================================================")
	router.Serve(a.Cfg.Server.Port)
}
