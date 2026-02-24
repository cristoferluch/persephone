package repository

import (
	"database/sql"
	_ "embed"
	"persephone/internal/entity"
)

//go:embed queries/indexes.sql
var queryIndexes string

type IndexRepository struct {
	conn *sql.DB
}

func NewIndexRepository(conn *sql.DB) *IndexRepository {
	return &IndexRepository{conn}
}

func (r *IndexRepository) GetIndexByTable(table string) ([]entity.Index, error) {
	rows, err := r.conn.Query(queryIndexes, table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	indexes := make([]entity.Index, 0)
	for rows.Next() {
		var index entity.Index
		if err := rows.Scan(
			&index.Name,
			&index.Description,
			&index.Keys,
		); err != nil {
			return nil, err
		}
		indexes = append(indexes, index)
	}
	return indexes, nil
}
