package entity

type Block struct {
	ID         int64  `db:"id"`
	PrevHash   string `db:"prevhash"`
	Data       string `db:"data"`
	Hash       string `db:"hash"`
	CreateTime int64
}

func (b *Block) TableName() string {
	return "blocks"
}
