package repository

import (
	"github.com/NoahOrberg/evileye/entity"
	"github.com/jmoiron/sqlx"
)

type Blocks interface {
	GetLatestBlock() (*entity.Block, error)
}

func NewBlocksRepository(db *sqlx.DB) Blocks {
	return &blocks{
		db: db,
	}
}

type blocks struct {
	db *sqlx.DB
}

func (b *blocks) GetLatestBlock() (*entity.Block, error) {
	block := new(entity.Block)
	row := b.db.QueryRow(`SELECT id, prevhash, data, create_time, hash
        FROM blocks
        WHERE create_time = (SELECT MAX(create_time) FROM blocks);`)
	if err := row.Scan(
		&block.Id,
		&block.Prevhash,
		&block.Data,
		&block.CreateTime,
		&block.Hash); err != nil {
		return nil, err
	}
	return block, nil
}
