package models

type URL struct {
	ID    string `json:"id"`
	Short string `json:"short"`
	Full  string `json:"full"`
}

type URLCreateRequest struct {
	Full string `json:"full"`
}
type URLUpdateRequest struct {
	Full string `json:"full"`
}
