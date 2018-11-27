package profile

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// This code is from:
// https://github.com/pitakill/rickandmortyapigowrapper/blob/master/http.go#L13
func makePetition(ops Options) (interface{}, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Can move on build options
	if len(ops.ID) > 0 {
		ops.Endpoint = ops.Endpoint + ops.ID
	}

	req, err := http.NewRequest(ops.Method, ops.Endpoint, bytes.NewBuffer(ops.Body))
	if err != nil {
		return nil, err
	}

	if len(ops.Token) > 0 {
		req.Header.Add("Authorization", ops.Token)
	}

	// Find by queryString
	if len(ops.Params) > 0 {
		q := req.URL.Query()
		for _, value := range ops.Params {
			q.Add(value.Field, value.Content)
		}
		req.URL.RawQuery = q.Encode()
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response interface{}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
