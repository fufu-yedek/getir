package records

import (
	"errors"
	apierrors "github.com/fufu-yedek/getir-challange/apihelper/errors"
	"github.com/fufu-yedek/getir-challange/apihelper/response"
	"github.com/fufu-yedek/getir-challange/gerrors"
	"github.com/sirupsen/logrus"
)

type Controller interface {
	ListRecords(params ListRecordParams) (response.Responder, error)
}

type controller struct {
	RecordRepository Repository
}

func (c controller) ListRecords(params ListRecordParams) (response.Responder, error) {
	// swagger:route POST /records Records List-Records
	//
	// List & filter records
	//
	//     Produces:
	//     - application/json
	//
	//     Consumes:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: listRecordsSwaggerResponse
	//       400: listRecordsSwaggerErrorResponse
	//       500: listRecordsSwaggerErrorResponse

	logger := logrus.WithFields(logrus.Fields{
		"location": "Controller - ListRecords",
		"params":   params.Body,
	})

	records, err := c.RecordRepository.FindWithCount(Filter{
		StartDate: params.Body.StartDate.ToTime(),
		EndDate:   params.Body.EndDate.ToTime(),
		MinCount:  params.Body.MinCount,
		MaxCount:  params.Body.MaxCount,
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

//NewDefaultController generates a new controller using MongoDB repository
func NewDefaultController() Controller {
	return NewController(NewDefaultMongoRepository())
}
