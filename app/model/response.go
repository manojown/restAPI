package model

type Response struct {
	Status  int         `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(status int, Message string, data interface{}) *Response {

	return &Response{
		Status:  status,
		Message: Message,
		Data:    data,
	}

}
