package domain

type Feed struct {
	ID          string `json:"id"`
	PostId      string `json:"post_id"`
	Username    string `json:"username"`
	CreatedTime string `json:"created_time"`
}
