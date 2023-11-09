package model

type Book struct {
	Id string `json:"id"`
    Title string `json:"title"`
    Author string `json:"author"`
    Issue_date string `json:"issue_date"`
    Publisher string `json:"publisher"`
    Description string `json:"description"`
    Tag string `json:"tag"`
    Last_update string `json:"last_update"`
    Last_update_date string `json:"last_update_date"`
}
