package profile

import (
	"github.com/mitchellh/mapstructure"
)

// GetUsers return list profile
func GetUsers(o Options) (*User, error) {
	data, err := makePetition(o)

	if err != nil {
		return &User{}, err
	}

	origins := new(User)

	if err := mapstructure.Decode(data, &origins); err != nil {
		return &User{}, err
	}

	return origins, nil
}

// GetProfiles return list profile
func GetProfiles(o Options) (*Profile, error) {
	data, err := makePetition(o)

	if err != nil {
		return &Profile{}, err
	}

	destinations := new(Profile)

	if err := mapstructure.Decode(data, &destinations); err != nil {
		return &Profile{}, err
	}

	return destinations, nil
}
