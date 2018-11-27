package profile

// Paginate struct
type Paginate struct {
	TotalCount int    `json:"totalCount"`
	PageCount  int    `json:"pageCount"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	SortBy     string `json:"sortBy"`
}

// Metadata struct
type Metadata struct {
	Paginate Paginate `json:"paginate"`
}

// Options struct
type Options struct {
	Endpoint string  `json:"endpoint"`
	Body     []byte  `json:"body"`
	Params   []Param `json:"params"`
	ID       string  `json:"id"`
	Token    string  `json:"token"`
	Method   string  `json:"method"`
}

// Param struct
type Param struct {
	Field   string `json:"field"`
	Content string `json:"content"`
}

// BodyAttributes struct
type BodyAttributes struct {
	Type       string                 `json:"type"`
	ID         string                 `json:"id"`
	Attributes map[string]interface{} `json:"attributes"` //TODO: Move to struct
}

// Body struct
type Body struct {
	Data BodyAttributes `json:"data"`
}
