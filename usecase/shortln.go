package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"test/amartha/router"

	"test/amartha/usecase/helper"
	"test/amartha/usecase/model"

	"github.com/julienschmidt/httprouter"
)

var helperRandString = helper.RandString

//GetShorten convert url to shortln
func (u *Usecase) GetShorten(w http.ResponseWriter, r *http.Request, ps httprouter.Params) router.Response {
	requestData := &model.ShortlnRequest{}
	err := json.NewDecoder(r.Body).Decode(requestData)
	if err != nil {
		return router.NewResponse().SetError(err, http.StatusBadRequest)
	}

	if requestData.Url == "" {
		return router.NewResponse().SetError(fmt.Errorf("url is not present"), http.StatusBadRequest)
	}

	resp := &model.ShortlnResponse{}
	if requestData.Shortcode == "" {
		newStr := ""
		isValid := false
		for !isValid {
			newStr = helperRandString(6)
			exist := u.DB.GetShortenByCode(newStr)
			if exist != nil {
				continue
			}

			isValid = true
		}
		requestData.Shortcode = newStr
	} else {
		//Validate
		ok, err := regexp.Match(helper.RegexpValidator, []byte(requestData.Shortcode))
		if err != nil || !ok {
			return router.NewResponse().SetError(fmt.Errorf("The shortcode fails to meet requirement"), http.StatusUnprocessableEntity)
		}

		exist := u.DB.GetShortenByCode(requestData.Shortcode)
		if exist != nil {
			return router.NewResponse().SetError(fmt.Errorf("The the desired shortcode is already in use"), http.StatusConflict)
		}

	}

	resp.Shortcode = requestData.Shortcode
	err = u.DB.CreateShortenCode(requestData)
	if err != nil {
		return router.NewResponse().SetError(err, http.StatusInternalServerError)
	}
	return router.NewResponse().SetData(resp).SetCode(http.StatusCreated)
}

func (u *Usecase) GetURL(w http.ResponseWriter, r *http.Request, ps httprouter.Params) router.Response {
	shrtn := ps.ByName("shorten")

	exist := u.DB.GetShortenByCode(shrtn)
	if exist == nil {
		return router.NewResponse().SetError(fmt.Errorf("The shortcode cannot be found in the system"), http.StatusNotFound)
	}

	err := u.DB.CountVisitingURL(shrtn)
	if err != nil {
		return router.NewResponse().SetError(err, http.StatusInternalServerError)
	}
	return router.NewResponse().SetCode(http.StatusFound).SetHeader("Location", exist.Url)

}

func (u *Usecase) GetURLStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) router.Response {
	shrtn := ps.ByName("shorten")

	exist := u.DB.GetShortenByCode(shrtn)
	if exist == nil {
		return router.NewResponse().SetError(fmt.Errorf("The shortcode cannot be found in the system"), http.StatusNotFound)
	}

	return router.NewResponse().SetData(exist)
}
