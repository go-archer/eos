package service

import (
    "context"
	"{{ .Project }}/internal/dto/request"
	"{{ .Project }}/internal/dto/response"
	"{{ .Project }}/internal/repository"
)

type {{ .FileName }}Service interface {
	Get{{ .FileName }}ById(ctx context.Context, in *request.{{ .FileName }}) (*response.{{ .FileName }}, error)
}

type {{ .TitleLower }}Service struct {
	*Service
	{{ .TitleLower }}Repository repository.{{ .FileName }}Repository
}

func New{{ .FileName }}Service(service *Service, {{ .TitleLower }}Repository repository.{{ .FileName }}Repository) {{ .FileName }}Service {
	return &{{ .TitleLower }}Service{
		Service:        service,
		{{ .TitleLower }}Repository: {{ .TitleLower }}Repository,
	}
}

func (s *{{ .TitleLower }}Service) Get{{ .FileName }}ById(ctx context.Context, in *request.{{ .FileName }}) (*response.{{ .FileName }}, error) {
     // TODO: business logic
	 // _, _ = s.{{ .TitleLower }}Repository.FirstById(ctx, in.ID)
	 return &response.{{ .FileName }}{}, nil
}
