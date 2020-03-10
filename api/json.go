package api

import (
	"bytes"
	"encoding/json"

	api "github.com/micro/go-micro/v2/api/proto"
)

// JSON JSON
func JSON(rsp *api.Response, code int32, i interface{}) error {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(i)

	if err != nil {
		return err
	}

	rsp.StatusCode = code
	rsp.Body = b.String()

	return nil
}
