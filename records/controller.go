package records

import (
	"errors"
	apierrors "github.com/fufuceng/getir-challange/apihelper/errors"
	"github.com/fufuceng/getir-challange/apihelper/response"
	"github.com/fufuceng/getir-challange/gerrors"
	"github.com/sirupsen/logrus"
)

type Controller interface {
	ListRecords(params ListRecordParams) (response.Responder, error)
}

type controller struct {
	RecordRepository Repository
}

func (c controller) ListRecords(params ListRecordParams) (response.Responder, error) {
	logger := logrus.WithFields(logrus.Fields{
		"location": "Controller - ListRecords",
		"params":   params,
	})

	records, err := c.RecordRepository.FindWithCount(Filter{
		StartDate: params.StartDate.ToTime(),
		EndDate:   params.EndDate.ToTime(),
		MinCount:  params.MinCount,
		MaxCount:  params.MaxCount,
	})

	if err != nil && !errors.Is(err, gerrors.ErrRecordNotFound) {
		logger.WithError(err).Error("error while finding records")
		return nil, apierrors.ErrInternalServer
	}

	return ListRecordsSerializer{Records: records}, nil
}

func NewController(recordRepository Repository) Controller {
	return controller{
		RecordRepository: recordRepository,
	}
}

func NewDefaultController() Controller {
	return NewController(NewDefaultMongoRepository())
}
