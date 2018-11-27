package profile

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestGestUsers --- Users
func TestGestUsers(t *testing.T) {
	ops := Options{
		Endpoint: EndpointUserAPI + "/users",
		Method:   "GET",
	}

	response, err := GetUsers(ops)

	if err != nil {
		t.Error(err)
	}

	data, err := ReadFile("test-data/users.json")

	if err != nil {
		t.Error(err)
	}

	result := new(User)

	json.Unmarshal(data, &result)

	comparation := cmp.Equal(result, response)

	if !comparation {
		t.Error("Response from GetUsers was:")
		t.Error(response)
		t.Error("The data against is being run this test is:")
		t.Error(result)
	}
}

func TestFindUserByEmail(t *testing.T) {
	attrs := make([]Param, 0)
	attrs = append(attrs, Param{Field: "email", Content: "abygromero@gmail.com"})

	ops := Options{
		Endpoint: EndpointUserAPI + "/users",
		Params:   attrs,
		Method:   "GET",
	}

	response, err := GetUsers(ops)

	if err != nil {
		t.Error(err)
	}

	data, err := ReadFile("test-data/users_filter_email.json")

	if err != nil {
		t.Error(err)
	}

	result := new(User)

	json.Unmarshal(data, &result)

	comparation := cmp.Equal(result, response)

	if !comparation {
		t.Error("Response from GetFindUserByEmail was:")
		t.Error(response)
		t.Error("The data against is being run this test is:")
		t.Error(result)
	}
}
