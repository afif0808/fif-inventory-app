package iServices

import "inventory/iModels"

type ErrorService interface {
	ErrorAndCode(error, int) iModels.IErrorModel
}
