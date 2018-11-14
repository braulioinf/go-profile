package profile

import (
	"encoding/json"
	"flag"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mitchellh/mapstructure"
)

var tokenFlag = flag.String("token", "", "Token needed for make the petition")

// TestGetProfileFirstPage --- Prfiles
func TestGetProfileFirstPage(t *testing.T) {

	if *tokenFlag == "" {
		t.Log("Token is required")
		os.Exit(0)
	}

	params := map[string]string{
		"page":   "1",
		"limit":  "10",
		"sortBy": "DESC",
	}

	attrs, err := buildArrayMaps(params)

	ops := Options{
		endpoint: "https://dev.api.culturacolectiva.com/profiles",
		params:   attrs,
		token:    *tokenFlag,
	}

	response, err := GetProfiles(ops)

	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/profiles_first_page_limit_10.json")

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

// TestFindProfileByEmail --- Prfiles
func TestFindProfileByEmail(t *testing.T) {
	params := map[string]string{
		"filter.email": "braulio.aguilar@culturacolectiva.com",
		"page":         "1",
		"limit":        "1",
		"sortBy":       "DESC",
	}

	attrs, err := buildArrayMaps(params)

	ops := Options{
		endpoint: "https://dev.api.culturacolectiva.com/profiles",
		params:   attrs,
		token:    *tokenFlag,
		method:   "GET",
	}

	response, err := GetProfiles(ops)

	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/profiles_filter_email.json")

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

// TestGestUsers --- Users
func TestGestUsers(t *testing.T) {
	ops := Options{
		endpoint: "http://localhost:3000/users",
		method:   "GET",
	}

	response, err := GetUsers(ops)

	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/users.json")

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
	params := map[string]string{
		"email": "u.marquez.j@gmail.com",
	}

	attrs, err := buildArrayMaps(params)

	ops := Options{
		endpoint: "http://localhost:3000/users",
		params:   attrs,
		method:   "GET",
	}

	response, err := GetUsers(ops)

	if err != nil {
		t.Error(err)
	}

	data, err := readFile("test-data/users_filter_email.json")

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

func TestPatchProfile(t *testing.T) {
	params := map[string]string{
		"filter.email": "braulio.aguilar@culturacolectiva.com",
		"page":         "1",
		"limit":        "1",
		"sortBy":       "DESC",
	}

	attrs, err := buildArrayMaps(params)
	getOps := Options{
		endpoint: "https://dev.api.culturacolectiva.com/profiles",
		params:   attrs,
		method:   "GET",
		token:    *tokenFlag,
	}

	response, err := GetProfiles(getOps)

	if len(response.Data) == 0 {
		t.Error("Don't found profile")
	}
	// End search by email

	// postOps
	data := map[string]interface{}{
		"data": map[string]interface{}{
			"type": "profiles",
			"attributes": map[string]interface{}{
				"alias":    "saimondev",
				"position": "develop",
			},
		},
	}

	bodyProfile, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}

	profileID := response.Data[0].ID
	postOps := Options{
		endpoint: "https://dev.api.culturacolectiva.com/profiles/" + profileID,
		token:    *tokenFlag,
		method:   "PATCH",
		body:     bodyProfile,
	}

	profileUpdated, err := makePetition(postOps)

	if err != nil {
		t.Error(err)
	}

	profileData := new(Data)

	if err := mapstructure.Decode(profileUpdated, &profileData); err != nil {
		t.Error(err)
	}

	profileFromJSON, err := readFile("test-data/profile_updated.json")

	profileJSONData := new(Data)

	json.Unmarshal(profileFromJSON, &profileJSONData)

	comparation := cmp.Equal(profileJSONData, profileData)

	if !comparation {
		t.Error("Response from PatchProfile was:")
		t.Error(profileData)
		t.Error("The data against is being run this test is:")
		t.Error(profileJSONData)
	}
}
