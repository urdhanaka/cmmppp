package dto

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewResponse(status bool, message string) Response {
	return Response{
		Status:  status,
		Message: message,
	}
}

func NewResponseWithData(status bool, message string, data any) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
