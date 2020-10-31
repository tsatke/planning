package server

type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Element interface{} `json:"element,omitempty"`
}
