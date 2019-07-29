package infrastructures

import (
	"database/sql"
	"fmt"
	"inventory/iInfrastructures"
)

type MySQLHandler struct {
	Conn *sql.DB
}

func (handler *MySQLHandler) Execute(query string, args ...interface{}) {
	handler.Conn.Exec(query, args...)
}

func (handler *MySQLHandler) Query(query string, args ...interface{}) (iInfrastructures.IRows, error) {
	rows, err := handler.Conn.Query(query, args...)

	if err != nil {
		fmt.Println(err)
		return new(MySqlRows), err
	}
	row := new(MySqlRows)
	row.Rows = rows

	return row, nil
}
func (handler *MySQLHandler) QuerySingle(query string, args ...interface{}) iInfrastructures.IRow {
	row := handler.Conn.QueryRow(query, args...)
	return row
}

// Multiple rows
type MySqlRows struct {
	Rows *sql.Rows
}

func (r MySqlRows) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return nil
}

func (r MySqlRows) Next() bool {
	return r.Rows.Next()
}

// DESC : Single row
type MySqlRowSingle struct {
	Row *sql.Row
}

func (r *MySqlRowSingle) Scan(dest ...interface{}) error {
	return r.Scan(dest...)
}
