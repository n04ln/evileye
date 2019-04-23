package entity

type Tarekomi struct {
	ID           int64  `db:"id"`
	Status       int64  `db:"status"`
	Threshold    int64  `db:"threshold"`
	TargetUserID int64  `db:"targetuserid"`
	URL          string `db:"url"`
	Description  string `db:"description"`
}

func (t *Tarekomi) TableName() string {
	return "tarekomi"
}
