package domain

type Post struct {
	ID          string `json:"id"`
	ProfileId   string `json:"profile_id"`
	Body        string `json:"body"`
	CreatedTime string `json:"created_time"`
}
