package redis

import (
	"context"
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
)

const (
	EthBlockCursorKey = "ethBlockCursor"
)

type redisCli struct {
	c *redis.Client
}

type RedisCli interface {
	// for: eth-->dip
	GetEthBlockCursor(ctx context.Context) (int64, error)
	SetEthBlockCursor(ctx context.Context, blockNumber int64)
	SaveEthTxidProcessReceiptOnDip(ctx context.Context, ethTxid, dipTxReceipt string)
	EthTxidExist(ctx context.Context, ethTxid string) bool

	// for: dip-->eth
}

func NewRedisCli(network string) RedisCli {
	cli := redis.NewClient(&redis.Options{
		Addr:     network,
		Password: "",
		DB:       0,
	})

	return &redisCli{c: cli}
}

func (c *redisCli) get(ctx context.Context, key string) string {
	val, err := c.c.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			log.Fatalf("do redis get failed:[%v],key:[%s]\n", err, key)
		}
	}

	return val
}

func (c *redisCli) set(ctx context.Context, key, value string) {
	err := c.c.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Fatalf("do redis set failed:[%v],key:[%s],value:[%s]\n", err, key, value)
	}
}

func (c *redisCli) GetEthBlockCursor(ctx context.Context) (int64, error) {
	val := c.get(ctx, EthBlockCursorKey)
	if val == "" {
		return 0, nil
	}

	tBlockCursor, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, err
	}

	return tBlockCursor, nil
}

func (c *redisCli) SetEthBlockCursor(ctx context.Context, blockNumber int64) {
	val := strconv.FormatInt(blockNumber, 10)
	c.set(ctx, EthBlockCursorKey, val)
}

func (c *redisCli) SaveEthTxidProcessReceiptOnDip(ctx context.Context, ethTxid, dipTxReceipt string) {
	c.set(ctx, ethTxid, dipTxReceipt)
}

func (c *redisCli) EthTxidExist(ctx context.Context, ethTxid string) bool {
	receipt := c.get(ctx, ethTxid)
	return len(receipt) > 0
}
