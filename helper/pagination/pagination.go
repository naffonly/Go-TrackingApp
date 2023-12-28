package pagination

type Pagination struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalItems int64 `json:"total_items"`
}

type QueryParam struct {
	Page  int
	Size  int
	Query string
}

func FormatPaginationResponse(message string, data any, pagination any) map[string]any {
	var responsePagination = map[string]any{}
	responsePagination["message"] = message
	responsePagination["data"] = data
	responsePagination["pagination"] = pagination
	return responsePagination
}

func FormatResponse(message string, data any, status int64) map[string]any {
	var responseResponse = map[string]any{}
	responseResponse["message"] = message
	responseResponse["data"] = data
	responseResponse["status"] = status
	return responseResponse
}
