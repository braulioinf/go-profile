package profile

// Options struct
type Options struct {
	endpoint string
	body     []byte
	params   []map[string]string
	id       string
	token    string
	method   string
}

// Profile struct
type Profile struct {
	Metadata struct {
		Paginate struct {
			TotalCount int    `json:"totalCount"`
			PageCount  int    `json:"pageCount"`
			Page       int    `json:"page"`
			Limit      int    `json:"limit"`
			SortBy     string `json:"sortBy"`
		}
	}
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes Attributes
	}
}

// Attributes struct - Profiles
type Attributes struct {
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

// User struct - Users
type User struct {
	Data []struct {
		Type       string `json:"type"`
		ID         int    `json:"id"`
		Attributes struct {
			ID                int    `json:"id"`
			Birthdate         string `json:"birthdate"`
			Photo             string `json:"photo"`
			Location          string `json:"location"`
			Phone             string `json:"phone"`
			Status            string `json:"status"`
			Greeting          string `json:"greeting"`
			Realm             string `json:"realm"`
			Username          string `json:"username"`
			Password          string `json:"password"`
			Email             string `json:"email"`
			EmailVerified     int    `json:"emailVerified"`
			VerificationToken string `json:"verificationToken"`
			Position          string `json:"position"`
			Description       string `json:"description"`
			Twitter           string `json:"twitter"`
			Facebook          string `json:"facebook"`
			Slug              string `json:"slug"`
			Type              string `json:"type"`
		}
	}
}

// Data send to profiles - PATCH METHOD
type Data struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes Attributes
}
