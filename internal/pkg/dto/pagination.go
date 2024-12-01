package dto

type PaginationDTO struct {
	sortColumn  string
	sortType    string
	LimitValue  int   `db:"limit_value"`
	OffsetValue int   `db:"offset_value"`
	IDs         []int `db:"ids"`
}

func NewPaginationDTO(sortColumn string, sortType string, limitValue int, offsetValue int, ids []int) *PaginationDTO {
	return &PaginationDTO{
		sortColumn:  sortColumn,
		sortType:    sortType,
		LimitValue:  limitValue,
		OffsetValue: offsetValue,
		IDs:         ids,
	}
}

func (dto *PaginationDTO) GetSortColumn() string {
	return dto.sortColumn
}

func (dto *PaginationDTO) GetSortType() string {
	return dto.sortType
}

func (dto *PaginationDTO) GetIDs() []int {
	return dto.IDs
}
