package profile

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestGetProfileSecondPage --- Profiles
func TestGetProfileSecondPage(t *testing.T) {

	if *tokenFlag == "" {
		t.Log("Token is required")
		os.Exit(0)
	}

	// Prepare filter
	profileAttrs := make([]Param, 0)
	profileAttrs = append(profileAttrs, Param{Field: "page", Content: "2"})
	profileAttrs = append(profileAttrs, Param{Field: "limit", Content: "10"})
	profileAttrs = append(profileAttrs, Param{Field: "sortBy", Content: "DESC"})

	ops := Options{
		Endpoint: EndpointProfileAPI,
		Params:   profileAttrs,
		Token:    *tokenFlag,
	}

	response, err := GetProfiles(ops)

	if err != nil {
		t.Error(err)
	}

	data, err := ReadFile("test-data/profiles_second_page_limit_10.json")

	if err != nil {
		t.Error(err)
	}

	result := new(Profile)

	json.Unmarshal(data, &result)

	comparation := cmp.Equal(result, response)

	if !comparation {
		t.Error("Response from GetProfileFirstPage was:")
		t.Error(response)
		t.Error("The data against is being run this test is:")
		t.Error(result)
	}
}

// TestFindProfileByEmail --- Profiles
func TestFindProfileByEmail(t *testing.T) {
	// Prepare filter
	profileAttrs := make([]Param, 0)
	profileAttrs = append(profileAttrs, Param{Field: "filter.email", Content: "elizabeth.flores@culturacolectiva.com"})
	profileAttrs = append(profileAttrs, Param{Field: "page", Content: "1"})
	profileAttrs = append(profileAttrs, Param{Field: "limit", Content: "1"})
	profileAttrs = append(profileAttrs, Param{Field: "sortBy", Content: "DESC"})

	ops := Options{
		Endpoint: EndpointProfileAPI,
		Params:   profileAttrs,
		Token:    *tokenFlag,
		Method:   "GET",
	}

	response, err := GetProfiles(ops)

	if err != nil {
		t.Error(err)
	}

	data, err := ReadFile("test-data/profiles_filter_email.json")

	if err != nil {
		t.Error(err)
	}

	result := new(Profile)

	json.Unmarshal(data, &result)

	comparation := cmp.Equal(result, response)

	if !comparation {
		t.Error("Response from GetFindProfileByEmail was:")
		t.Error(response)
		t.Error("The data against is being run this test is:")
		t.Error(result)
	}
}
