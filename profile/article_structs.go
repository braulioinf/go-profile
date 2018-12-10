package profile

// Author struct
type Author struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Type        string `json:"type"`
	Image       string `json:"image"`
	Position    string `json:"position"`
	Description string `json:"description"`
	Facebook    string `json:"facebook"`
	Twitter     string `json:"twitter"`
}

// Seo struct
type Seo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	Keywords    []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
}

// Editor struct
type Editor struct {
}

// Meta struct
type Meta struct {
	Tags []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	}
	References []struct {
		ID    string `json:"id"`
		Title string `json:"title,omitempty"`
		URL   string `json:"url,omitempty"`
	}
}

// Facebook struct
type Facebook struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Ads struct
type Ads struct {
	IsCommercial bool `json:"isCommercial"`
}

// ArticleData struct
type ArticleData struct {
	Type       string            `json:"type"`
	ID         string            `json:"id"`
	Attributes ArticleAttributes `json:"attributes"`
}

// ArticleAttributes struct
type ArticleAttributes struct {
	LegacyID    string   `json:"legacyId,omitempty"`
	Title       string   `json:"title"`
	Subtitle    string   `json:"subtitle"`
	Content     string   `json:"content"`
	Category    string   `json:"category"`
	Type        string   `json:"type"`
	Language    string   `json:"language"`
	Status      string   `json:"status"`
	Author      Author   `json:"author"`
	Editor      Editor   `json:"editor"`
	Seo         Seo      `json:"seo"`
	Ads         Ads      `json:"ads"`
	Facebook    Facebook `json:"facebook"`
	Meta        Meta     `json:"meta"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
	PublishedAt string   `json:"publishedAt"`
}

// Article struct
type Article struct {
	Metadata Metadata      `json:"metadata"`
	Data     []ArticleData `json:"data"`
}
