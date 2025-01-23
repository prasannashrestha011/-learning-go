package dtos

type ResponseDto struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	StatusCode int         `json:"status_code"`
}
type DataDto struct {
	Data       interface{} `json:"data"`
	StatusCode int         `json:"status_code"`
	Success    bool        `json:"success"`
	Message    *string     `json:"message,omitempty"`
}
