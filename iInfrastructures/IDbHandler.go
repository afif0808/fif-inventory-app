package iInfrastructures

type IDbHandler interface {
	Execute(query string, args ...interface{})
	Query(query string, args ...interface{}) (IRows, error)
	QuerySingle(query string, args ...interface{}) IRow
}

type IRows interface {
	Scan(dest ...interface{}) error
	Next() bool
}
type IRow interface {
	Scan(dest ...interface{}) error
}
