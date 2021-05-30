package memrecords

import (
	"errors"
	apierrors "github.com/fufuceng/getir-challange/apihelper/errors"
	"github.com/fufuceng/getir-challange/apihelper/response"
	"github.com/fufuceng/getir-challange/gerrors"
	"github.com/sirupsen/logrus"
)

type Controller interface {
	CreateOrUpdate(params CreateOrUpdateParams) (response.Responder, error)
	Retrieve(params RetrieveParams) (response.Responder, error)
}

type controller struct {
	repository Repository
}

func (c controller) CreateOrUpdate(params CreateOrUpdateParams) (response.Responder, error) {
	logger := logrus.WithFields(logrus.Fields{
		"location": "Controller - CreateOrUpdate",
		"params":   params,
	})

	record, err := c.repository.CreateOrUpdate(Record{
		Key:   params.Key,
		Value: params.Value,
	})

	if err != nil {
		logger.WithError(err).Error("error while creating new record")
		return nil, apierrors.ErrInternalServer
	}

	return RecordSerializer{Record: record}, nil
}

func (c controller) Retrieve(params RetrieveParams) (response.Responder, error) {
	logger := logrus.WithFields(logrus.Fields{
		"location": "Controller - Retrieve",
		"params":   params,
	})

	record, err := c.repository.FindOne(Filter{
		Key: params.Key,
	})

	if err != nil {
		if errors.Is(err, gerrors.ErrRecordNotFound) {
			logrus.WithError(err).Warn("could not find record")
			return nil, apierrors.NewUserReadableErrf("could not find record")
		}

		logger.WithError(err).Error("error while finding record")
		return nil, apierrors.ErrInternalServer
	}

	return RecordSerializer{Record: record}, nil
}

func NewController(repository Repository) Controller {
	return controller{repository: repository}
}

func NewDefaultController() Controller {
	return NewController(NewDefaultInMemRepository())
}
