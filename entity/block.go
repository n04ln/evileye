package entity

type Block struct {
	Id         int64
	Prevhash   string
	Data       string
	Hash       string
	CreateTime int64
}

func (b Block) TableName() string {
	return "blocks"
}
