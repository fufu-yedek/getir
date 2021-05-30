package records

import (
	apierrors "github.com/fufuceng/getir-challange/apihelper/errors"
	"github.com/fufuceng/getir-challange/apihelper/request"
	"github.com/fufuceng/getir-challange/apihelper/response"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Router interface {
	ListRecords(w http.ResponseWriter, req *http.Request)
	Register(mux *http.ServeMux)
}

type router struct {
	RecordController Controller
}

func (r router) Register(mux *http.ServeMux) {
	mux.HandleFunc("/records", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			r.ListRecords(w, req)
		}
	})
}

func (r router) ListRecords(w http.ResponseWriter, req *http.Request) {
	logger := logrus.WithField("location", "Router - ListRecords")

	var params ListRecordParams
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

	resp, err := r.RecordController.ListRecords(params)
	response.GenerateResponse(w, resp, err)
}

func NewRouter(rc Controller) Router {
	return router{
		RecordController: rc,
	}
}

func NewDefaultRouter() Router {
	return NewRouter(NewDefaultController())
}
