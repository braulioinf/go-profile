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
	if len(ops.id) > 0 {
		ops.endpoint = ops.endpoint + ops.id
	}

	req, err := http.NewRequest(ops.method, ops.endpoint, bytes.NewBuffer(ops.body))
	if err != nil {
		return nil, err
	}

	if len(ops.token) > 0 {
		req.Header.Add("Authorization", ops.token)
	}

	// Find by queryString
	if ops.params != nil {
		q := req.URL.Query()
		for _, value := range ops.params {
			for i, v := range value {
				q.Add(i, v)
			}
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
