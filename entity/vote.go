package entity

type Vote struct {
	ID          int64  `db:"id"`
	UserID      int64  `db:"int64"`
	TarekomiID  int64  `db:"tarekomiid"`
	Description string `db:"description"`
}

func (v *Vote) TableName() string {
	return "votes"
}
