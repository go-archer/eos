package request

type {{ .FileName }} struct {
	ID    int64  `json:"id" binding:"required"`
}

