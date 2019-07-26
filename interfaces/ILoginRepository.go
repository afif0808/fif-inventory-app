package interfaces

import "inventory/modelInterfaces"

type ILoginRepository interface {
	GetLoginByToken(token string) (login modelInterfaces.ILoginModel, err error)
}
