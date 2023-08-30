package response

type {{ .FileName }} struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
}

type {{ .FileName }}List struct {
	Items []*Area `json:"items"`
}
