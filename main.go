package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Cultura-Colectiva-Tech/go-profile/profile"
	"github.com/fatih/color"
)

const (
	urlPrefix         = "https://"
	urlSuffix         = ".culturacolectiva.com/"
	urlProfilesSuffix = "profiles"
	email             = "email"
	slug              = "slug"
	urlArticlesSuffix = "articles"
)

var (
	userEmailFlag    *string
	profileEmailFlag *string
	userSlugFlag     *string
	tokenFlag        *string
	environmentFlag  *string
	dataUserFlag     *string
	api              string
	green            = color.New(color.FgGreen).SprintFunc()
	red              = color.New(color.FgRed).SprintFunc()
	blue             = color.New(color.FgHiBlue).SprintFunc()
	selectedUser     int
	taskFlag         *string
	limitFlag        *string
	pageFlag         *string
	typePostFlag     *string
	startDateFlag    *string
	endDateFlag      *string
	statusPostFlag   *string
	authorSlugFlag   *string
	authorIDFlag     *string
)

func main() {
	tokenFlag = flag.String("token", "", "Token needed to make the petition")
	environmentFlag = flag.String("environment", "dev", "Environment to make metition {dev, staging, prod}")
	dataUserFlag = flag.String("data-users", "http://localhost:3000", "URL for get users")
	userEmailFlag = flag.String("user-email", "", "Email user to get source info")
	userSlugFlag = flag.String("user-slug", "", "Slug user to get source info")
	profileEmailFlag = flag.String("profile-email", "", "Email profile to push info")
	taskFlag = flag.String("task", "profiles", "Task name to action, {profiles, articles}")
	// Migrate
	limitFlag = flag.String("limit", "50", "Limit of items in the response")
	pageFlag = flag.String("page", "1", "Number of the page where start")
	typePostFlag = flag.String("type-post", "POST", "Article type to be searched {VIDEO,POST}. Default: video")
	startDateFlag = flag.String("start-date", "2018-01-01", "Year to bring Article, Default: 2018-01-01")
	endDateFlag = flag.String("end-date", "2018-12-31", "Month to bring Articles. Default: 2018-12-31")
	statusPostFlag = flag.String("status-post", "STATUS_PUBLISHED", "Article status to be searched. Default: published")

	// Special filters
	authorSlugFlag = flag.String("author-slug", "", "Author's Slug (slug from profile)")
	authorIDFlag = flag.String("author-id", "", "Author's ID (providerId from profile)")

	envs := map[string]string{
		"dev":     "dev.api",
		"prod":    "api-v2",
		"staging": "staging.api",
	}

	flag.Parse()

	api = envs[*environmentFlag]

	options := []string{"profiles", "articles"}

	if !profile.Contains(options, *taskFlag) {
		fmt.Println("Task not valid")
		os.Exit(0)
	}

	if *taskFlag == urlProfilesSuffix {
		taskMigrateProfile()
		os.Exit(0)
	}

	if *taskFlag == urlArticlesSuffix {

		if *authorSlugFlag != "" && *authorIDFlag != "" {
			fmt.Printf("Can't use same time 'author-slug' and 'author-id' params.")
			os.Exit(0)
		}

		taskFillMetadata()
	}
}
