package helpers

type MetaResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Response struct {
	Meta    MetaResponse `json:"meta"`
	Data    interface{}  `json:"data"`
	Success bool         `json:"success"`
}

func ResponseSuccess(data interface{}, message string, code int) Response {
	return Response{
		Meta: MetaResponse{
			Code:    code,
			Status:  "success",
			Message: message,
		},
		Data:    data,
		Success: true,
	}
}
