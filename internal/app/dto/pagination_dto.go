package dto

type PaginationDTO struct {
	sortColumn  string
	sortType    string
	limitValue  int
	offsetValue int
}

func NewPaginationDTO(sortColumn string, sortType string, limitValue int, offsetValue int) *PaginationDTO {
	return &PaginationDTO{
		sortColumn:  sortColumn,
		sortType:    sortType,
		limitValue:  limitValue,
		offsetValue: offsetValue,
	}
}

func (dto *PaginationDTO) GetSortColumn() string {
	return dto.sortColumn
}

func (dto *PaginationDTO) GetSortType() string {
	return dto.sortType
}

func (dto *PaginationDTO) GetLimitValue() int {
	return dto.limitValue
}

func (dto *PaginationDTO) GetOffsetValue() int {
	return dto.offsetValue
}
