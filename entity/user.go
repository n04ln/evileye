package entity

type User struct {
	ID         int64  `db:"id"`
	ScreenName string `db:"screenname"`
	// Password   []byte `db:"password"`
	Password string `db:"password"`
}

func (u *User) TableName() string {
	return "users"
}
