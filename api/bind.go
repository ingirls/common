package api

import (
	"encoding/json"

	api "github.com/micro/go-micro/v2/api/proto"
)

// Bind Bind
func Bind(req *api.Request, i interface{}) error {
	if err := json.Unmarshal([]byte(req.Body), &i); err != nil {
		return err
	}
	return nil
}
