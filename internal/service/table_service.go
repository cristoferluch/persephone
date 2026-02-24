package service

import (
	"persephone/internal/entity"
	"persephone/internal/repository"
)

type TableService struct {
	tableRepository  *repository.TableRepository
	columnRepository *repository.ColumnRepository
	indexRepository  *repository.IndexRepository

	cachedTables      []entity.Table
	cachedTableIndex  map[string][]entity.Index
	cachedTableColumn map[string][]entity.Column
}

func NewTableService(
	tableRepository *repository.TableRepository,
	columnRepository *repository.ColumnRepository,
	indexRepository *repository.IndexRepository,
) *TableService {
	return &TableService{
		tableRepository:   tableRepository,
		columnRepository:  columnRepository,
		indexRepository:   indexRepository,
		cachedTables:      make([]entity.Table, 0),
		cachedTableIndex:  make(map[string][]entity.Index),
		cachedTableColumn: make(map[string][]entity.Column),
	}
}

func (s *TableService) GetColumnByTable(tableName string) ([]entity.Column, error) {

	if columns, ok := s.cachedTableColumn[tableName]; ok {
		return columns, nil
	}

	columns, err := s.columnRepository.GetColumnByTable(tableName)
	if err != nil {
		return nil, err
	}
	s.cachedTableColumn[tableName] = columns

	return columns, nil
}

func (s *TableService) GetIndexByTable(tableName string) ([]entity.Index, error) {

	if indexes, ok := s.cachedTableIndex[tableName]; ok {
		return indexes, nil
	}

	indexes, err := s.indexRepository.GetIndexByTable(tableName)
	if err != nil {
		return nil, err
	}
	s.cachedTableIndex[tableName] = indexes

	return indexes, nil
}

func (s *TableService) FindAll(search string) ([]entity.Table, error) {

	if len(s.cachedTables) > 0 {
		return s.cachedTables, nil
	}

	table, err := s.tableRepository.FindAll(search)
	if err != nil {
		return nil, err
	}
	s.cachedTables = table

	return table, nil
}
