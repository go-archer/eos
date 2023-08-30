package repository

import (
    "context"
	"{{ .Project }}/internal/model"
)

type {{ .FileName }}Repository interface {
	FirstById(ctx context.Context, id int64) (*model.{{ .FileName }}, error)
}

type {{ .TitleLower }}Repository struct {
	*Repository
}

func New{{ .FileName }}Repository(repository *Repository) {{ .FileName }}Repository {
	return &{{ .TitleLower }}Repository{
		Repository: repository,
	}
}

func (r *{{ .TitleLower }}Repository) FirstById(ctx context.Context, id int64) (*model.{{ .FileName }}, error) {
	var {{ .TitleLower }} model.{{ .FileName }}
	// TODO: query db
	return &{{ .TitleLower }}, nil
}
