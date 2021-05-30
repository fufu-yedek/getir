package memrecords

import (
	apierrors "github.com/fufuceng/getir-challange/apihelper/errors"
	"github.com/fufuceng/getir-challange/apihelper/request"
	"github.com/fufuceng/getir-challange/apihelper/response"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Router interface {
	CreateOrUpdate(w http.ResponseWriter, req *http.Request)
	Retrieve(w http.ResponseWriter, req *http.Request)
	Register(mux *http.ServeMux)
}

type router struct {
	Controller Controller
}

func (r router) CreateOrUpdate(w http.ResponseWriter, req *http.Request) {
	logger := logrus.WithField("location", "Memrecord Router - CreateOrUpdate")

	var params CreateOrUpdateParams
	if err := request.ParseJSON(req, &params.Body); err != nil {
		logger.WithError(err).Error("error while parsing request")
		response.GenerateResponse(w, nil, apierrors.ErrInternalServer)
		return
	}

	if err := params.Validate(); err != nil {
		logger.WithError(err).Error("error while validating parameters")
		response.GenerateResponse(w, nil, err)
		return
	}

	resp, err := r.Controller.CreateOrUpdate(params)
	response.GenerateResponse(w, resp, err)
}

func (r router) Retrieve(w http.ResponseWriter, req *http.Request) {
	logger := logrus.WithField("location", "Memrecord Router - Retrieve")

	var params RetrieveParams
	if err := request.ParseQuery(req, &params); err != nil {
		logger.WithError(err).Error("error while parsing request")
		response.GenerateResponse(w, nil, apierrors.ErrInternalServer)
		return
	}

	if err := params.Validate(); err != nil {
		logger.WithError(err).Error("error while validating parameters")
		response.GenerateResponse(w, nil, err)
		return
	}

	resp, err := r.Controller.Retrieve(params)
	response.GenerateResponse(w, resp, err)
}

func (r router) Register(mux *http.ServeMux) {
	mux.HandleFunc("/in-memory", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			r.CreateOrUpdate(w, req)
		case http.MethodGet:
			r.Retrieve(w, req)
		}
	})
}

func NewRouter(controller Controller) Router {
	return router{Controller: controller}
}

func NewDefaultRouter() Router {
	return NewRouter(NewDefaultController())
}
