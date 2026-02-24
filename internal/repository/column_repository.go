package repository

import (
	"database/sql"
	_ "embed"
	"persephone/internal/entity"
)

//go:embed queries/columns.sql
var queryColumns string

type ColumnRepository struct {
	conn *sql.DB
}

func NewColumnRepository(conn *sql.DB) *ColumnRepository {
	return &ColumnRepository{conn}
}

func (r *ColumnRepository) GetColumnByTable(table string) ([]entity.Column, error) {
	rows, err := r.conn.Query(queryColumns, table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	indexes := make([]entity.Column, 0)
	for rows.Next() {
		var column entity.Column
		err = rows.Scan(
			&column.Name,
			&column.Type,
			&column.Length,
			&column.Precision,
			&column.IsNullable,
			&column.HasPrimaryKey,
		)
		if err != nil {
			return nil, err
		}
		indexes = append(indexes, column)

	}
	return indexes, nil
}
