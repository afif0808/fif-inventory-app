package iModels

type IErrorModel interface {
	Error() string
	ErrorCode() int
}
