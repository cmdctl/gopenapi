package v3_0

type Pullrequest struct {
	Author     User       `json:"author"`
	Id         int        `json:"id"`
	Repository Repository `json:"repository"`
	Title      string     `json:"title"`
}

type Repository struct {
	Owner User   `json:"owner"`
	Slug  string `json:"slug"`
}

type User struct {
	Username string `json:"username"`
	Uuid     string `json:"uuid"`
}
