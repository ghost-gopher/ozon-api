package utils

type List struct {
	Items  interface{} `json:"items,omitempty"`
	LastId string      `json:"last_id,omitempty"`
	Total  int32       `json:"total,omitempty"`
}
