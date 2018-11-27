package profile

// UserAttributes bla bla
type UserAttributes struct {
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

// User struct - Users
type User struct {
	Data []UserData `json:"data"`
}

// UserData bla bla
type UserData struct {
	Type       string         `json:"type"`
	ID         int            `json:"id"`
	Attributes UserAttributes `json:"attributes"`
}
