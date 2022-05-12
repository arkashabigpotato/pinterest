package models

type User struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	IsAdmin    bool   `json:"is_admin"`
	BirthDate  string `json:"birth_date"`
	Username   string `json:"user_name"`
	ProfileImg string `json:"profile_img"`
	Status     string `json:"status"`
}

type Message struct {
	ID           int    `json:"id"`
	FromID       int    `json:"from_id"`
	ToID         int    `json:"to_id"`
	Text         string `json:"text"`
	DateTime     string `json:"date_time"`
	FromUsername string `json:"from_username"`
	ToUsername   string `json:"to_username"`
	IsFromMe     bool   `json:"is_from_me"`
}

type Pin struct {
	ID            int    `json:"id"`
	Description   string `json:"description"`
	LikesCount    int    `json:"likes_count"`
	DislikesCount int    `json:"dislikes_count"`
	AuthorID      int    `json:"author_id"`
	PinLink       string `json:"pin_link"`
}

type SavedPin struct {
	PinID  int `json:"pin_id"`
	UserID int `json:"user_id"`
}

type Comment struct {
	ID        int    `json:"id"`
	IsDeleted bool   `json:"is_deleted"`
	PinID     int    `json:"pin_id"`
	Text      string `json:"text"`
	AuthorID  int    `json:"author_id"`
	DateTime  string `json:"date_time"`
	Username  string `json:"user_name"`
}

const dateTimeFormat = "2006-01-02 15:04:05"
