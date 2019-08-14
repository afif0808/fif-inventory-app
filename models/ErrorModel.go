package models

type ErrorModel struct {
	Err     error
	ErrCode int
}

func (err ErrorModel) Error() string {
	return err.Err.Error()
}
func (err ErrorModel) ErrorCode() int {
	return err.ErrCode
}
