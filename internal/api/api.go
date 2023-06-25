package api

import (
	"encoding/json"
)

type Response struct {
	Data json.RawMessage `json:"data,omitempty"`
}
