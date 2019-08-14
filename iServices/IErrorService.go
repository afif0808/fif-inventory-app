package iServices

import "inventory/iModels"

type IErrorService interface {
	ErrorAndCode(error, int) iModels.IErrorModel
}
