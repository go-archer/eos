package response

type {{ .FileName }} struct {
	ID         int64   `json:"id"`
}

type {{ .FileName }}List struct {
	Items []*{{ .FileName }} `json:"items"`
}
