package profile

// ProfileAttributes bla bla
type ProfileAttributes struct {
	ProviderID  string `json:"providerId"`
	Name        string `json:"name"`
	Alias       string `json:"alias"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	Birthday    string `json:"birthday"`
	Status      string `json:"status"`
	Role        string `json:"role"`
	Picture     string `json:"picture"`
	Slug        string `json:"slug"`
	Position    string `json:"position"`
	Description string `json:"description"`
	Facebook    string `json:"facebook"`
	Twitter     string `json:"twitter"`
}

// Profile struct
type Profile struct {
	Metadata Metadata      `json:"metadata"`
	Data     []ProfileData `json:"data"`
}

// ProfileData struct
type ProfileData struct {
	Type       string            `json:"type"`
	ID         string            `json:"id"`
	Attributes ProfileAttributes `json:"attributes"`
}
