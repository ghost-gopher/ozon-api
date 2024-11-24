package utils

type Sort struct {
	Filter interface{} `json:"filter,omitempty"`
	LastId string      `json:"last_id,omitempty"`
	Limit  int64       `json:"limit,omitempty"`
}
