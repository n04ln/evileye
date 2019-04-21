package models

type Star struct {
	ID         int64 `db:"id"`
	UserID     int64 `db:"userid"`
	TarekomiID int64 `db:"tarekomiid"`
}

func (s *Star) TableName() string {
	return "stars"
}
