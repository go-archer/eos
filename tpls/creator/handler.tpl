package handler

import (
	"{{ .Project }}/internal/dto/request"
	"{{ .Project }}/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type {{ .FileName }}Handler interface {
	Get{{ .FileName }}ById(ctx *gin.Context)
	Update{{ .FileName }}(ctx *gin.Context)
}

type {{ .TitleLower }}Handler struct {
	*Handler
	{{ .TitleLower }}Service service.{{ .FileName }}Service
}

func New{{ .FileName }}Handler(handler *Handler, {{ .TitleLower }}Service service.{{ .FileName }}Service) {{ .FileName }}Handler {
	return &{{ .TitleLower }}Handler{
		Handler:     handler,
		{{ .TitleLower }}Service: {{ .TitleLower }}Service,
	}
}

func (h {{ .TitleLower }}Handler) Get{{ .FileName }}ById(ctx *gin.Context) {
	in := &request.{{ .FileName }}{}

	if err := h.Bind(ctx, in); err!=nil {
	    h.Error(ctx,err)
	    return
	}

	{{ .TitleLower }}, err := h.{{ .TitleLower }}Service.Get{{ .FileName }}ById(ctx, in)
	h.log.Info("Get{{ .FileName }}ByID", zap.Any("{{ .TitleLower }}", {{ .TitleLower }}))
	if err != nil {
		h.Error(ctx, err)
		return
	}
	h.Success(ctx, {{ .TitleLower }})
}

func (h {{ .TitleLower }}Handler) Update{{ .FileName }}(ctx *gin.Context) {
	h.Success(ctx, nil)
}
