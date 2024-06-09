package paginate

import (
	"gorm.io/gorm"
)

type Page[T any] struct {
	Data []T  `json:"data"`
	Meta Meta `json:"meta"`
}

type Meta struct {
	CurrentPage int `json:"current_page"`
	From        int `json:"from"`
	PerPage     int `json:"per_page"`
	To          int `json:"to"`
}

func Simple[T any](query *gorm.DB, request Request) (*Page[T], error) {
	models := make([]T, 0, request.PageSize)

	err := query.Offset((request.Page - 1) * request.PageSize).Limit(request.PageSize).Find(&models).Error
	if err != nil {
		return nil, err
	}

	from := (request.Page - 1) * request.PageSize
	return &Page[T]{
		Data: models,
		Meta: Meta{
			CurrentPage: request.Page,
			From:        from + 1,
			PerPage:     request.PageSize,
			To:          from + len(models),
		},
	}, nil
}
