package repository

import (
	"context"
	"time"

	"github.com/NoahOrberg/evileye/entity"
	"github.com/NoahOrberg/evileye/log"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var ch = make(chan struct{}, 1)

type Blocks interface {
	GetLatestBlock() (*entity.Block, error)
	InsertBlock(ctx context.Context, data, prevHash, hash string) (*entity.Block, error)
}

func NewBlocksRepository(db *sqlx.DB) Blocks {
	ch <- struct{}{}
	return &blocks{
		db:           db,
		insertedHash: make([]string, 2, 2),
	}
}

type blocks struct {
	db           *sqlx.DB
	insertedHash []string
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

	<-ch
	defer func() {
		ch <- struct{}{}
	}()

	var bl *entity.Block
	var err error
	if bl, err = b.GetLatestBlock(); err != nil {
		return nil, err
	}
	if bl.PrevHash == prevHash {
		log.L().Info("BLOCKING CREATE BLOCK",
			zap.String("data", data),
			zap.String("prevHash", prevHash),
			zap.String("hash", hash))
		return nil, nil
	}
	if bl.Hash != prevHash {
		log.L().Info("CANNOT CREATE BLOCK COZ INVALID INTEGRITY",
			zap.String("data", data),
			zap.String("actual_prevHash", bl.Hash),
			zap.String("expected_prevHash", prevHash),
			zap.String("hash", hash))
		return nil, nil
	}
	if bl.CreateTime-1 <= time.Now().Unix() && time.Now().Unix() <= bl.CreateTime+1 {
		log.L().Info("CANNOT CREATE BLOCK COZ INVALID INTEGRITY (TIME LAG)",
			zap.String("data", data),
			zap.String("actual_prevHash", bl.Hash),
			zap.String("expected_prevHash", prevHash),
			zap.String("hash", hash))
		return nil, nil
	}

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
