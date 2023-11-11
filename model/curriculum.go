package model

type Curriculum struct {
    Id string `json:"id"`
    Title string `json:"title"`
    Chapter string `json:"chapter"`
    Url string `json:"url"`
    Description string `json:"description"`
    Tag string `json:"tag"`
    Last_update string `json:"last_update"`
    Last_update_date string `json:"last_update_date"`
}