package infrastructures

import (
	"database/sql"
	"fmt"
	"inventory/interfaces"
)

type MySQLHandler struct {
	Conn *sql.DB
}

func (handler *MySQLHandler) Execute(query string, args ...interface{}) {
	handler.Conn.Exec(query, args...)
}

func (handler *MySQLHandler) Query(query string, args ...interface{}) (interfaces.IRow, error) {
	rows, err := handler.Conn.Query(query, args...)

	if err != nil {
		fmt.Println(err)
		return new(MySqlRow), err
	}
	row := new(MySqlRow)
	row.Rows = rows

	return row, nil
}

type MySqlRow struct {
	Rows *sql.Rows
}

func (r MySqlRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return nil
}

func (r MySqlRow) Next() bool {
	return r.Rows.Next()
}
