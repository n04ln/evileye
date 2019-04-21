package models

type User struct {
	ID         int64  `db:"id"`
	ScreenName string `db:"screenname"`
	Password   []byte `db:"password"`
}

func (u *User) TableName() string {
	return "users"
}
