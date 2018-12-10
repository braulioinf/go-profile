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

// SetProfile func
func SetProfile(o Options) (interface{}, error) {
	data, err := makePetition(o)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetArticles func
func GetArticles(o Options) (*Article, error) {
	data, err := makePetition(o)

	if err != nil {
		return &Article{}, err
	}

	articles := new(Article)

	if err := mapstructure.Decode(data, &articles); err != nil {
		return &Article{}, err
	}

	return articles, nil
}

// SetArticle func
func SetArticle(o Options) (interface{}, error) {
	data, err := makePetition(o)

	if err != nil {
		return nil, err
	}

	return data, nil
}
