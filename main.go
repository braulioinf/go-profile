package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Cultura-Colectiva-Tech/go-profile/profile"

	"github.com/fatih/color"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

const (
	urlPrefix         = "https://"
	urlSuffix         = ".culturacolectiva.com/"
	urlProfilesSuffix = "profiles"
)

var (
	emailUserFlag    *string
	emailProfileFlag *string
	tokenFlag        *string
	envFlag          *string
	dataUserFlag     *string
	api              string
	green            = color.New(color.FgGreen).SprintFunc()
	red              = color.New(color.FgRed).SprintFunc()
)

func main() {
	emailUserFlag = flag.String("email-user", "", "Email user to get source info")
	emailProfileFlag = flag.String("email-profile", "", "Email profile to push info")
	tokenFlag = flag.String("token", "", "Token needed to make the petition")
	envFlag = flag.String("env", "dev", "Environment to make metition {dev, staging, prod}")
	dataUserFlag = flag.String("data-users", "http://localhost:3000", "URL for get users")

	envs := map[string]string{
		"dev":     "dev.api",
		"prod":    "api-v2",
		"staging": "staging.api",
	}

	flag.Parse()

	api = envs[*envFlag]

	if api != "" {
		migrateProfile()
	}
}

// MigrateProfile func
func migrateProfile() {
	profileURL := urlPrefix + api + urlSuffix + urlProfilesSuffix

	fmt.Printf("Working profile in %s\n", green(profileURL))
	fmt.Printf("Working user in %s\n", green(*dataUserFlag))

	userAttrs := make([]profile.Param, 0)
	userAttrs = append(userAttrs, profile.Param{Field: "email", Content: *emailUserFlag})

	userOps := profile.Options{
		Endpoint: *dataUserFlag + "/users",
		Params:   userAttrs,
		Method:   "GET",
	}

	fmt.Printf("Searching in user: %s\n", green(*emailUserFlag))
	userData, err := profile.GetUsers(userOps)
	if err != nil {
		log.Println(red(err))
	}

	// Prepare to search Profile
	profileAttrs := make([]profile.Param, 0)
	profileAttrs = append(profileAttrs, profile.Param{Field: "filter.email", Content: *emailProfileFlag})
	profileAttrs = append(profileAttrs, profile.Param{Field: "page", Content: "1"})
	profileAttrs = append(profileAttrs, profile.Param{Field: "limit", Content: "1"})
	profileAttrs = append(profileAttrs, profile.Param{Field: "sortBy", Content: "DESC"})

	profileOps := profile.Options{
		Endpoint: profileURL,
		Params:   profileAttrs,
		Method:   "GET",
		Token:    *tokenFlag,
	}

	fmt.Printf("Searching in profile: %s\n", green(*emailProfileFlag))
	profileData, err := profile.GetProfiles(profileOps)
	if err != nil {
		log.Println(red(err))
	}

	if len(userData.Data) == 0 || len(profileData.Data) == 0 {
		log.Println("User or Profile is empty")
		os.Exit(0)
	}

	// Available fields to set
	responseAttrs := structs.Map(userData.Data[0].Attributes)
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

	fmt.Printf("Profile for email: %s was updated\n", *emailProfileFlag)
}
