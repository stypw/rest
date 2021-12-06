package orm

import (
	"database/sql"
)

type field struct {
	FieldType  string
	FieldName  string
	FieldValue interface{}
}

type Orm struct {
	Db        *sql.DB
	TableName string
}

func makeFields(rows *sql.Rows) ([]*field, []interface{}, error) {
	if cts, err := rows.ColumnTypes(); err == nil {
		fields := make([]*field, 0)
		pointers := make([]interface{}, 0)
		for _, ct := range cts {
			fd := &field{FieldType: ct.DatabaseTypeName(), FieldName: ct.Name()}
			switch ct.DatabaseTypeName() {
			case "INT", "BIGINT":
				{
					var v int64
					fd.FieldValue = &v
					pointers = append(pointers, &v)
				}
			case "VARCHAR", "TEXT", "NVARCHAR":
				{
					var v string
					fd.FieldValue = &v
					pointers = append(pointers, &v)
				}
			case "DECIMAL":
				{
					var v float64
					fd.FieldValue = &v
					pointers = append(pointers, &v)
				}
			case "BOOL":
				{
					var v bool
					fd.FieldValue = &v
					pointers = append(pointers, &v)
				}
			}
			fields = append(fields, fd)
		}
		return fields, pointers, nil
	} else {
		return nil, nil, err
	}
}
