package cache

import (
	"iris-cn/model"
	"iris-cn/repositories"
	"time"

	"github.com/goburrow/cache"
	"github.com/mlogclub/simple"
)

var UserTokenCache = newUserTokenCache()

type userTokenCache struct {
	cache cache.LoadingCache
}

func newUserTokenCache() *userTokenCache {
	return &userTokenCache{
		cache: cache.NewLoadingCache(
			func(key cache.Key) (value cache.Value, e error) {
				value = repositories.UserTokenRepository.GetByToken(simple.DB(), key.(string))
				return
			},
			cache.WithMaximumSize(1000),
			cache.WithExpireAfterAccess(60*time.Minute),
		),
	}
}

func (c *userTokenCache) Get(token string) *model.UserToken {
	if len(token) == 0 {
		return nil
	}
	val, err := c.cache.Get(token)
	if err != nil {
		return nil
	}
	if val != nil {
		return val.(*model.UserToken)
	}
	return nil
}

func (c *userTokenCache) Invalidate(token string) {
	c.cache.Invalidate(token)
}
