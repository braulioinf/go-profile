package profile

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/fatih/structs"
	"github.com/google/go-cmp/cmp"
	"github.com/mitchellh/mapstructure"
)

var (
	userEmail    string
	profileEmail string
	emailTest    string
)

func TestMigrateProfileEmail(t *testing.T) {
	t.Run("When search user with an email valid then response a data user", func(t *testing.T) {
		t.Helper()
		emailExpected := "patricia.cordero@culturacolectiva.com"
		emailTest = "patricia.cordero@culturacolectiva.com"

		// Prepare attributes to search
		userParams := make([]Param, 0)
		userParams = append(userParams, Param{Field: "email", Content: emailTest})

		userOps := Options{
			Endpoint: EndpointUserAPI + "/users",
			Params:   userParams,
			Method:   "GET",
		}

		userData, err := GetUsers(userOps)
		if err != nil {
			t.Error(err)
		}

		email := userData.Data[0].Attributes.Email

		if emailExpected != email {
			t.Error("Response from GetUsers was:")
			t.Error(email)
			t.Error("The data against is being run this test is:")
			t.Error(emailTest)
		}
	})

	t.Run("When search profile with an email valid then response an data profile", func(t *testing.T) {
		t.Helper()
		emailExpected := "patricia.cordero@culturacolectiva.com"
		emailTest = "patricia.cordero@culturacolectiva.com"

		profileAttrs := make([]Param, 0)
		profileAttrs = append(profileAttrs, Param{Field: "filter.email", Content: emailTest})
		profileAttrs = append(profileAttrs, Param{Field: "page", Content: "1"})
		profileAttrs = append(profileAttrs, Param{Field: "limit", Content: "1"})
		profileAttrs = append(profileAttrs, Param{Field: "sortBy", Content: "DESC"})

		profileOps := Options{
			Endpoint: EndpointProfileAPI,
			Params:   profileAttrs,
			Method:   "GET",
			Token:    *tokenFlag,
		}

		profileData, err := GetProfiles(profileOps)
		if err != nil {
			t.Error(err)
		}

		email := profileData.Data[0].Attributes.Email

		if emailExpected != email {
			t.Error("Response from GetProfiles was:")
			t.Error(email)
			t.Error("The data against is being run this test is:")
			t.Error(emailTest)
		}
	})

	t.Run("When I have an user valid and profile valid then response with profile updated", func(t *testing.T) {
		emailTest := "patricia.cordero@culturacolectiva.com"

		// Prepare to search User
		userAttrs := make([]Param, 0)
		userAttrs = append(userAttrs, Param{Field: "email", Content: emailTest})

		userOps := Options{
			Endpoint: EndpointUserAPI + "/users",
			Params:   userAttrs,
			Method:   "GET",
		}

		userData, err := GetUsers(userOps)
		if err != nil {
			t.Error(err)
		}

		// Prepare to search Profile
		profileAttrs := make([]Param, 0)
		profileAttrs = append(profileAttrs, Param{Field: "filter.email", Content: emailTest})
		profileAttrs = append(profileAttrs, Param{Field: "page", Content: "1"})
		profileAttrs = append(profileAttrs, Param{Field: "limit", Content: "1"})
		profileAttrs = append(profileAttrs, Param{Field: "sortBy", Content: "DESC"})

		profileOps := Options{
			Endpoint: EndpointProfileAPI,
			Params:   profileAttrs,
			Method:   "GET",
			Token:    *tokenFlag,
		}

		profileData, err := GetProfiles(profileOps)
		if err != nil {
			t.Error(err)
		}

		if len(userData.Data) == 0 || len(profileData.Data) == 0 {
			t.Error("User or Profile is empty")
		}

		// Available fields to set
		responseAttrs := structs.Map(userData.Data[0].Attributes)
		available := []string{"Username", "Birthday", "Slug", "Position", "Description", "Facebook", "Twitter"}

		patchAttrs := make(map[string]interface{})
		for field, content := range responseAttrs {
			if Contains(available, field) {
				patchAttrs[strings.ToLower(field)] = content
			}
		}

		profileID := profileData.Data[0].ID
		data := Body{
			Data: BodyAttributes{
				Type:       "profiles",
				ID:         profileID,
				Attributes: patchAttrs,
			},
		}

		bodyProfile, err := json.Marshal(data)
		if err != nil {
			t.Error(err)
		}

		patchOps := Options{
			Endpoint: EndpointProfileAPI + "/" + profileID,
			Token:    *tokenFlag,
			Method:   "PATCH",
			Body:     bodyProfile,
		}

		profileUpdated, err := makePetition(patchOps)
		if err != nil {
			t.Error(err)
		}

		profileResponse := new(ProfileData)

		if err := mapstructure.Decode(profileUpdated, &profileResponse); err != nil {
			t.Error(err)
		}

		// Read local json
		profileFromJSON, err := ReadFile("test-data/profile_migrated_one.json")
		if err != nil {
			t.Error(err)
		}

		profileJSONData := new(ProfileData)
		json.Unmarshal(profileFromJSON, &profileJSONData)

		comparation := cmp.Equal(profileJSONData, profileResponse)

		if !comparation {
			t.Error("Response from PathProfile was:")
			t.Error(profileResponse)
			t.Error("The data against is being run this test is:")
			t.Error(profileJSONData)
		}
	})
}
