package entity

type SuccessResponse struct {
	Error   bool        `json:"error" example:"false"`
	Message string      `json:"message" example:"Create Data Succeed"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Error   bool   `json:"error" example:"true"`
	Message string `json:"message" example:"Create Data Failed"`
}
