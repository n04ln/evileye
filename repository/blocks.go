package repository

import (
	"context"
	"time"

	"github.com/NoahOrberg/evileye/entity"
	"github.com/NoahOrberg/evileye/log"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Blocks interface {
	GetLatestBlock() (*entity.Block, error)
	InsertBlock(ctx context.Context, data, prevHash, hash string) (*entity.Block, error)
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
		&block.ID,
		&block.PrevHash,
		&block.Data,
		&block.CreateTime,
		&block.Hash); err != nil {
		return nil, err
	}
	return block, nil
}

func (b *blocks) InsertBlock(ctx context.Context,
	data, prevHash, hash string) (*entity.Block, error) {
	block := new(entity.Block)

	block.Data = data
	block.PrevHash = prevHash
	block.Hash = hash
	block.CreateTime = time.Now().Unix()

	res, err := b.db.ExecContext(ctx,
		`INSERT INTO blocks(data, prevhash, create_time, hash) VALUES (?, ?, ?, ?)`,
		block.Data, block.PrevHash, block.CreateTime, block.Hash)
	if err != nil {
		log.L().Error("ExecContext in InsertBlock is failed",
			zap.Error(err))
		return nil, err
	}

	block.ID, err = res.LastInsertId()
	if err != nil {
		log.L().Error("InsertBlock response's LastInsertedID is failed",
			zap.Error(err))
		return nil, err
	}

	return block, nil
}
