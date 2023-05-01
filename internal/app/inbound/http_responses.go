package inbound

import "encoding/json"

type HttpError struct {
	Error   error  `json:"erro"`
	Message string `json:"message"`
}

func (e HttpError) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Error   string `json:"erro"`
		Message string `json:"message"`
	}{
		Error:   e.Error.Error(),
		Message: e.Message,
	})
}

type HttpSucess struct {
	Message string `json:"message"`
}
