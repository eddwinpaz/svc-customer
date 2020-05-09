package entity

// Response http body as JSON struct
type Response struct {
	Status      bool        `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

// CustomerResponse response body
type CustomerResponse struct {
	Customer *Customer `json:"customer"`
}

// Pagination Displays current Page, total rows and pages available from a Query SELECT Returns
type Pagination struct {
	Pages       int `json:"pages"`
	CurrentPage int `json:"current_page"`
	TotalRows   int `json:"total_rows"`
}

// CustomerPaginationResponse Deliver Custom Response for Fetch()
type CustomerPaginationResponse struct {
	Customer   []*Customer `json:"customers"`
	Pagination Pagination  `json:"pagination"`
}
