package profile

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	userSlug string
)

func TestSearchUserBySlug(t *testing.T) {
	slugTest := "carolina-romero"

	userParams := make([]Param, 0)
	userParams = append(userParams, Param{Field: "slug", Content: slugTest})

	userOps := Options{
		Endpoint: EndpointUserAPI + "/users",
		Params:   userParams,
		Method:   "GET",
	}

	userData, err := GetUsers(userOps)
	if err != nil {
		t.Error(err)
	}

	if len(userData.Data) == 0 {
		t.Error("User don't exist")
	}

	t.Run("When search user with a Slug valid then response two results", func(t *testing.T) {
		t.Helper()

		// Read local json
		userFromJSON, err := ReadFile("test-data/users_filter_slug.json")
		if err != nil {
			t.Error(err)
		}

		result := new(User)
		json.Unmarshal(userFromJSON, &result)
		comparation := cmp.Equal(result, userData)

		if !comparation {
			t.Error("Response from GetFindUserBySlug was:")
			t.Error(userData)
			t.Error("The data against is being run this test is:")
			t.Error(result)
		}
	})

	t.Run("When I have two user results then display in terminal two options (Second option is nice)", func(t *testing.T) {
		t.Helper()
		slugExpected := "carolina-romero"

		for _, v := range userData.Data {
			fmt.Printf("Deploying slug: %s for email: %s\n", v.Attributes.Slug, v.Attributes.Email)
		}

		comparation := cmp.Equal(slugExpected, userData.Data[1].Attributes.Slug)

		if !comparation {
			t.Error("Response for second result was:")
			t.Error(userData.Data[1].Attributes.Slug)
			t.Error("The data against is being run this test is:")
			t.Error(slugExpected)
		}
	})
}
