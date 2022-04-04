package cache

import (
	"github.com/VictoriaMetrics/fastcache"

	"wb_L0/internal/config"
)

type (
	cacheMem struct {
		*fastcache.Cache
	}
)

func NewCache(cfg *config.Config) Cache {
	return &cacheMem{fastcache.New(cfg.Cache.Size)}
}

func (c *cacheMem) Set(k []byte, v []byte) {
	c.Cache.Set(k, v)
}

func (c *cacheMem) Get(dst []byte, k []byte) []byte {
	return c.Cache.Get(dst, k)
}
