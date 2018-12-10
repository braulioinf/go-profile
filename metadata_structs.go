package main

import "github.com/Cultura-Colectiva-Tech/go-profile/profile"

// Paginate struct
type Paginate struct {
	Page       int `json:"page"`
	PageCount  int `json:"pageCount"`
	TotalCount int `json:"totalCount"`
}

// Types func
type Types struct {
	Type      string                `json:"type"`
	Data      []profile.ArticleData `json:"data"`
	Available []string              `json:"available"`
}
