package usecase

import (
	"net/http"
	"test/amartha/config"
	"test/amartha/database"
	"test/amartha/router"

	"github.com/julienschmidt/httprouter"
)

type Usecase struct {
	Cfg config.Config
	DB  database.DBInterface
}

type Interactor interface {
	GetShorten(w http.ResponseWriter, r *http.Request, ps httprouter.Params) router.Response
	GetURL(w http.ResponseWriter, r *http.Request, ps httprouter.Params) router.Response
	GetURLStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) router.Response
}
