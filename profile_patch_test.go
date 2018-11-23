package profile

import (
	"encoding/json"
	"flag"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mitchellh/mapstructure"
)

var tokenFlag = flag.String("token", "", "Token needed for make the petition")

func TestPatchProfile(t *testing.T) {

	profileAttrs := make([]Param, 0)
	profileAttrs = append(profileAttrs, Param{field: "filter.email", content: "abygromero@gmail.com"})
	profileAttrs = append(profileAttrs, Param{field: "page", content: "1"})
	profileAttrs = append(profileAttrs, Param{field: "limit", content: "1"})
	profileAttrs = append(profileAttrs, Param{field: "sortBy", content: "DESC"})

	getOps := Options{
		Endpoint: endpointProfileAPI,
		Params:   profileAttrs,
		Method:   "GET",
		Token:    *tokenFlag,
	}

	response, err := GetProfiles(getOps)

	if len(response.Data) == 0 {
		t.Error("Don't found profile")
	}

	// postOps
	data := map[string]interface{}{
		"data": map[string]interface{}{
			"type": "profiles",
			"attributes": map[string]interface{}{
				"name":  "Abril Romero",
				"email": "abygromero@gmail.com",
			},
		},
	}

	bodyProfile, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}

	profileID := response.Data[0].ID
	postOps := Options{
		Endpoint: endpointProfileAPI + "/" + profileID,
		Token:    *tokenFlag,
		Method:   "PATCH",
		Body:     bodyProfile,
	}

	profileUpdated, err := makePetition(postOps)

	if err != nil {
		t.Error(err)
	}

	profileData := new(ProfileData)

	if err := mapstructure.Decode(profileUpdated, &profileData); err != nil {
		t.Error(err)
	}

	profileFromJSON, err := readFile("test-data/profile_updated.json")

	profileJSONData := new(ProfileData)

	json.Unmarshal(profileFromJSON, &profileJSONData)

	comparation := cmp.Equal(profileJSONData, profileData)

	if !comparation {
		t.Error("Response from PatchProfile was:")
		t.Error(profileData)
		t.Error("The data against is being run this test is:")
		t.Error(profileJSONData)
	}
}
