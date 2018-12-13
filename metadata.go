package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Cultura-Colectiva-Tech/go-profile/profile"
)

var (
	forArtMigrated  = []string{"Id", "Name", "Type", "Image", "Position", "Description", "Facebook", "Twitter"}
	forArtOriginals = []string{"Name", "Slug", "Type", "Image", "Position", "Description", "Facebook", "Twitter"}
)

func taskFillMetadata() {
	articleURL := urlPrefix + api + urlSuffix + urlArticlesSuffix

	params, _ := fillParams(*pageFlag)

	options := profile.Options{
		Endpoint: articleURL,
		Params:   params,
		Method:   "GET",
		Token:    *tokenFlag,
	}

	articlesResponse, err := profile.GetArticles(options)
	if err != nil {
		fmt.Println(red(err))
	}

	if len(articlesResponse.Data) == 0 {
		fmt.Println("Article list is empty, try with other filters")
		os.Exit(0)
	}

	actual, pageCount, _ := getPaginate(articlesResponse)

	for actual <= pageCount {
		page := strconv.FormatInt(int64(actual), 10)
		// Get articles by block
		params2, _ := fillParams(page)

		options2 := profile.Options{
			Endpoint: articleURL,
			Params:   params2,
			Method:   "GET",
			Token:    *tokenFlag,
		}

		response, err := profile.GetArticles(options2)
		if err != nil {
			fmt.Println(red(err))
		}

		pageInt, pageCount, totalCount := getPaginate(response)

		pageLog := strconv.Itoa(pageInt)
		pageCountLog := strconv.Itoa(pageCount)
		totalCountLog := strconv.Itoa(totalCount)

		fmt.Printf("Processing: Page %s of %s with %s total items\n", green(pageLog), green(pageCountLog), green(totalCountLog))

		originals := []profile.ArticleData{} // Filter by author => providerId
		migrated := []profile.ArticleData{}  // Filter by author => slug

		for _, a := range response.Data {
			if len(a.Attributes.LegacyID) == 0 {
				originals = append(originals, a)
			} else {
				migrated = append(migrated, a)
			}
		}

		allTypes := make([]Types, 0)

		// Build slice
		allTypes = append(allTypes, Types{
			Type:      "providerId",
			Data:      originals,
			Available: forArtOriginals,
		})

		allTypes = append(allTypes, Types{
			Type:      "slug",
			Data:      migrated,
			Available: forArtMigrated,
		})

		// Execute generic func for migrated, new articles
		for _, types := range allTypes {
			fmt.Printf("......... Init search filter for %s in %d items ......... \n", green(types.Type), len(types.Data))
			items, err := searchProfileFromArticle(types.Data, types.Type)
			if err != nil {
				fmt.Printf(red(err))
				continue
			}

			executePatch(items, types.Type, types.Available)
			fmt.Printf("......... Finished filter for %s ......... \n\n", green(types.Type))
		}

		allTypes = nil
		actual++
	}
}
