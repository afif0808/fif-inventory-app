package services

import (
	"inventory/iModels"
	"inventory/models"
)

type ErrorService struct {
}

func (service *ErrorService) ErrorAndCode(err error, errCode int) iModels.IErrorModel {
	return models.ErrorModel{Err: err, ErrCode: errCode}
}
