package dtos

type ErrorResponse struct {
	Error string `json:"error" example:"Something went wrong"`
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
}
