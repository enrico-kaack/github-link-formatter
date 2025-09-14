package github

type GhResponse struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	State  string `json:"state"`
}
