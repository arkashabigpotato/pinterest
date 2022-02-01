package models

type User struct {
	ID 		  int
	Email 	  string
	Password  string
	IsAdmin   bool
	BirthDate string    `json:"birth_date"`
	Username  string    `json:"user_name"`
}

type Message struct {
	ID 		 int        `json:"id"`
	FromID   int        `json:"from_id"`
	ToID 	 int        `json:"to_id"`
	Text 	 string     `json:"text"`
	DateTime string     `json:"date_time"`
}

type Pin struct {
	ID 			  int
	Description   string `json:"description"`
	LikesCount 	  int    `json:"likes_count"`
	DislikesCount int    `json:"dislikes_count"`
	AuthorID 	  int    `json:"author_id"`
	PinLink 	  string `json:"pin_link"`
}

type SavedPin struct {
	PinID  int       `json:"pin_id"`
	UserID int       `json:"user_id"`
}

type Comment struct {
	ID 		  int
	IsDeleted bool
	PinID 	  int
	Text 	  string
	AuthorID  int
	DateTime  string
}

const dateTimeFormat = "2006-01-02 15:04:05"