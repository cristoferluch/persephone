package repository

import (
	"database/sql"
	_ "embed"
	"persephone/internal/entity"
)

//go:embed queries/tables.sql
var queryTables string

type TableRepository struct {
	conn *sql.DB
}

func NewTableRepository(conn *sql.DB) *TableRepository {
	return &TableRepository{conn}
}

func (r *TableRepository) FindAll(search string) ([]entity.Table, error) {
	rows, err := r.conn.Query(queryTables, "%"+search+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tables := make([]entity.Table, 0)
	for rows.Next() {
		var table entity.Table
		if err := rows.Scan(&table.Name); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}

	return tables, nil
}
