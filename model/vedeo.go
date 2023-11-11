package model

type Vedeo struct {
    Id string `json:"id"`
    Title string `json:"title"`
    Author string `json:"author"`
    Url string `json:"url"`
    Description string `json:"description"`
    Tag string `json:"tag"`
    Last_update string `json:"last_update"`
    Last_update_date string `json:"last_update_date"`
}