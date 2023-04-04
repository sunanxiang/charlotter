package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var GlobalCache *cache.Cache

func Init() {
	GlobalCache = cache.New(time.Minute*5, time.Hour*24)
}
