package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/Cultura-Colectiva-Tech/go-profile/profile"
	"github.com/fatih/structs"
)

// Build data to PATCH articles
func executePatch(items map[string]interface{}, t string, available []string) {
	articleURL := urlPrefix + api + urlSuffix + urlArticlesSuffix

	attrs := make(map[string]interface{})

	for articleKey, profileItem := range items {

		patchNewAttrs := make(map[string]interface{})
		fmt.Printf("Process PATCH article ID %s ==> ", green(articleKey))

		props := structs.Map(profileItem)

		for key, val := range props {
			// Only articles migrated don't exist provideID
			if key == "ProviderID" {
				patchNewAttrs["id"] = val
				continue
			}

			if profile.Contains(available, key) {
				if len(val.(string)) > 0 {
					patchNewAttrs[strings.ToLower(key)] = val
				}
			}
		}

		attrs["author"] = patchNewAttrs
		data := profile.Body{
			Data: profile.BodyAttributes{
				Type:       "articles",
				ID:         string(articleKey),
				Attributes: attrs,
			},
		}

		body, err := json.Marshal(data)
		if err != nil {
			fmt.Println(red(err))
			continue
		}

		patchOps := profile.Options{
			Endpoint: articleURL + "/" + string(articleKey),
			Token:    *tokenFlag,
			Method:   "PATCH",
			Body:     body,
		}

		_, error := profile.SetArticle(patchOps)
		if error != nil {
			fmt.Printf("merge failed with error: %s", red(error))
			continue
		}

		fmt.Println("profile data was merged into author.")
	}
}

func fillParams(page string) ([]profile.Param, error) {
	params := make([]profile.Param, 0)
	params = append(params, profile.Param{Field: "filter.type", Content: *typePostFlag})
	params = append(params, profile.Param{Field: "filter.status", Content: *statusPostFlag})
	params = append(params, profile.Param{Field: "filter.startDate", Content: *startDateFlag})
	params = append(params, profile.Param{Field: "filter.endDate", Content: *endDateFlag})
	params = append(params, profile.Param{Field: "page", Content: page})
	params = append(params, profile.Param{Field: "limit", Content: *limitFlag})
	params = append(params, profile.Param{Field: "sortBy", Content: "DESC"})

	if len(*authorSlugFlag) > 0 {
		params = append(params, profile.Param{Field: "filter.authorSlug", Content: *authorSlugFlag})
	}

	if len(*authorIDFlag) > 0 {
		params = append(params, profile.Param{Field: "filter.authorId", Content: *authorIDFlag})
	}

	return params, nil
}

// Parse pagination
func getPaginate(a *profile.Article) (page, pageCount, totalCount int) {

	p := Paginate{
		Page:       a.Metadata.Paginate.Page,
		PageCount:  a.Metadata.Paginate.PageCount,
		TotalCount: a.Metadata.Paginate.TotalCount,
	}

	page = p.Page
	pageCount = p.PageCount
	totalCount = p.TotalCount
	return
}

// Search prolife entity from author article metadata
func searchProfileFromArticle(articles []profile.ArticleData, genre string) (map[string]interface{}, error) {
	profileURL := urlPrefix + api + urlSuffix + urlProfilesSuffix

	response := make(map[string]interface{}, 0)
	cache := make(map[string]interface{}, 0)

	content := ""

	// Articles
	for d, article := range articles {

		if genre == "slug" {
			content = article.Attributes.Author.Slug
		} else {
			content = article.Attributes.Author.ID
		}

		fmt.Printf("%d._ Searching in article ID: %s with %s = %s ", d+1, blue(string(article.ID)), blue(genre), blue(content))

		_, cached := cache[content]
		if cached {
			fmt.Printf(".......... %s => Found profile ID from CACHE.\n", green(genre))
			response[article.ID] = cache[content]
			continue
		}

		// Prepare to search Profile
		param := make([]profile.Param, 0)
		param = append(param, profile.Param{Field: "filter." + genre, Content: content})
		param = append(param, profile.Param{Field: "page", Content: "1"})
		param = append(param, profile.Param{Field: "limit", Content: "1"})
		param = append(param, profile.Param{Field: "sortBy", Content: "DESC"})

		options := profile.Options{
			Endpoint: profileURL,
			Params:   param,
			Method:   "GET",
			Token:    *tokenFlag,
		}

		profileResponse, err := profile.GetProfiles(options)
		if err != nil {
			log.Println(red(err))
			continue
		}

		// Don't exist profile
		if len(profileResponse.Data) == 0 {
			fmt.Println(red("......... profile don't exist"))
			continue
		}

		pro := profileResponse.Data[0]

		// temp
		cache[content] = pro.Attributes

		fmt.Printf(".......... %s => Found profile ID: %s. \n", green(genre), green(pro.ID))

		response[article.ID] = pro.Attributes
	}

	return response, nil
}
