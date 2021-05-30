package response

import (
	"encoding/json"
	"errors"
	apierrors "github.com/fufuceng/getir-challange/apihelper/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

var DefaultSuccessResponse = BaseResponse{
	Code: 0,
	Msg:  "Success",
}

type Responder interface {
	Response() interface{}
}

type BaseResponse struct {
	// example: 0
	Code uint `json:"code"`
	// example: Success
	Msg string `json:"msg"`
}

func GenerateResponse(w http.ResponseWriter, resp Responder, dErr error) {
	logger := logrus.WithField("location", "GenerateResponse")

	w.Header().Set(ContentTypeHeader, ContentTypeApplicationJson)

	var (
		body interface{}
		code int
	)

	if resp != nil {
		body = resp.Response()
		code = http.StatusOK
	} else {
		if errors.Is(dErr, apierrors.ErrUserReadable) {
			code = apierrors.CodeUserReadableError
			body = BaseResponse{
				Code: apierrors.CodeUserReadableError,
				Msg:  dErr.Error(),
			}
		} else if errors.Is(dErr, apierrors.ErrInternalServer) {
			code = apierrors.CodeInternalServerError
			body = BaseResponse{
				Code: apierrors.CodeInternalServerError,
				Msg:  dErr.Error(),
			}
		} else {
			code = apierrors.CodeInternalServerError
			body = BaseResponse{
				Code: apierrors.CodeInternalServerError,
				Msg:  "internal server error",
			}

			logger.WithError(dErr).Error("unknown error detected")
		}
	}

	writeJSON(w, body, code)
}

func writeJSON(w http.ResponseWriter, body interface{}, code int) {
	logger := logrus.WithField("location", "writeJSON")

	w.WriteHeader(code)

	r, err := json.Marshal(body)
	if err != nil {
		logger.WithError(err).Error("error while marshaling json")
		return
	}

	if _, err := w.Write(r); err != nil {
		logger.WithError(err).Error("error while writing body")
		return
	}

	return
}
