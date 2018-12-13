package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Cultura-Colectiva-Tech/go-profile/profile"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

// MigrateProfile func
func taskMigrateProfile() {
	profileURL := urlPrefix + api + urlSuffix + urlProfilesSuffix

	userKey := 0

	fmt.Printf("Working profile in %s\n", green(profileURL))
	fmt.Printf("Working user in %s\n", green(*dataUserFlag))

	content := *userEmailFlag
	field := email

	if len(*userSlugFlag) > 0 {
		content = *userSlugFlag
		field = slug
	}

	fmt.Printf("Searching user by %s\n", green(field))

	userAttrs := make([]profile.Param, 0)
	userAttrs = append(userAttrs, profile.Param{Field: field, Content: content})

	userOps := profile.Options{
		Endpoint: *dataUserFlag + "/users",
		Params:   userAttrs,
		Method:   "GET",
	}

	fmt.Printf("Searching in user: %s\n", green(content))
	userData, err := profile.GetUsers(userOps)
	if err != nil {
		log.Println(red(err))
	}

	// Prepare to search Profile
	profileAttrs := make([]profile.Param, 0)
	profileAttrs = append(profileAttrs, profile.Param{Field: "filter.email", Content: *profileEmailFlag})
	profileAttrs = append(profileAttrs, profile.Param{Field: "page", Content: "1"})
	profileAttrs = append(profileAttrs, profile.Param{Field: "limit", Content: "1"})
	profileAttrs = append(profileAttrs, profile.Param{Field: "sortBy", Content: "DESC"})

	profileOps := profile.Options{
		Endpoint: profileURL,
		Params:   profileAttrs,
		Method:   "GET",
		Token:    *tokenFlag,
	}

	fmt.Printf("Searching in profile: %s\n", green(*profileEmailFlag))
	profileData, err := profile.GetProfiles(profileOps)
	if err != nil {
		log.Println(red(err))
	}

	if len(userData.Data) == 0 || len(profileData.Data) == 0 {
		log.Println("User or Profile is empty")
		os.Exit(0)
	}

	if len(userData.Data) > 1 {
		fmt.Printf("%s items found\n\n", green(len(userData.Data)))

		for k, v := range userData.Data {
			fmt.Printf("%s) slug: %s with email: %s\n", green(k+1), green(v.Attributes.Slug), green(v.Attributes.Email))
		}

		_, err := fmt.Scanf("%d", &selectedUser)
		if err != nil {
			fmt.Printf("%s\n", err)
		}

		userKey = selectedUser - 1
	}

	// Available fields to set
	responseAttrs := structs.Map(userData.Data[userKey].Attributes)
	available := []string{"Username", "Birthday", "Slug", "Position", "Description", "Facebook", "Twitter"}

	patchAttrs := make(map[string]interface{})
	for field, content := range responseAttrs {
		if profile.Contains(available, field) {
			patchAttrs[strings.ToLower(field)] = content
			fmt.Printf("Building body to patch: [%s] = %s\n", green(strings.ToLower(field)), content)
		}
	}

	profileID := profileData.Data[0].ID
	data := profile.Body{
		Data: profile.BodyAttributes{
			Type:       "profiles",
			ID:         profileID,
			Attributes: patchAttrs,
		},
	}

	bodyProfile, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	patchOps := profile.Options{
		Endpoint: profileURL + "/" + profileID,
		Token:    *tokenFlag,
		Method:   "PATCH",
		Body:     bodyProfile,
	}

	response, err := profile.SetProfile(patchOps)
	if err != nil {
		log.Println(red(err))
	}

	profileResponse := new(profile.ProfileData)

	if err := mapstructure.Decode(response, &profileResponse); err != nil {
		log.Println(err)
	}

	fmt.Printf("Profile for email: %s was updated\n", *profileEmailFlag)
}
