package interfaces

type IDbHandler interface {
	Execute(query string, args ...interface{})
	Query(query string, args ...interface{}) (IRow, error)
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
}
