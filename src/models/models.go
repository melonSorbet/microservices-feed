package models

import "time"

type Post []struct {
	PostID    string    `json:"post_id"`
	Caption   string    `json:"caption"`
	Content   string    `json:"content"`
	Username  string    `json:"username"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Media     []struct {
		MediaID          string `json:"MediaId"`
		FileURL          string `json:"FileUrl"`
		Width            int    `json:"Width"`
		Height           int    `json:"Height"`
		FileType         string `json:"FileType"`
		OriginalFileName string `json:"OriginalFileName"`
	} `json:"media"`
}

type Users struct {
	Following []struct {
		UserID           string    `json:"user_id"`
		Username         string    `json:"username"`
		Email            string    `json:"email"`
		DisplayName      string    `json:"display_name"`
		BirthDate        string    `json:"birth_date"`
		Location         string    `json:"location"`
		Bio              string    `json:"bio"`
		ProfilePicURL    string    `json:"profile_pic_url"`
		ProfilePicWidth  int       `json:"profile_pic_width"`
		ProfilePicHeight int       `json:"profile_pic_height"`
		CreatedAt        time.Time `json:"created_at"`
	} `json:"following"`
	Total int `json:"total"`
}
