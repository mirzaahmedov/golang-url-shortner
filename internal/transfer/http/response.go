package http

type MetaInfo struct {
	TotalCount int
}
type ResponseBody struct {
	data any
	meta MetaInfo
}

type ErrorResponse struct {
}
