package cache

import (
	"iris-cn/model"
	"iris-cn/model/constants"
	"iris-cn/repositories"
	"time"

	"github.com/goburrow/cache"
	"github.com/gqzcl/simple"
)

var (
	articleRecommendCacheKey = "recommend_articles_cache"
	articleHotCacheKey       = "hot_articles_cache"
)

var ArticleCache = newArticleCache()

type articleCache struct {
	recommendCache cache.LoadingCache
	hotCache       cache.LoadingCache
}

func newArticleCache() *articleCache {
	return &articleCache{
		recommendCache: cache.NewLoadingCache(
			func(key cache.Key) (value cache.Value, e error) {
				value = repositories.ArticleRepository.Find(simple.DB(),
					simple.NewSqlCnd().Where("status = ?", constants.StatusOk).Desc("id").Limit(50))
				return
			},
			cache.WithMaximumSize(1),
			cache.WithRefreshAfterWrite(30*time.Minute),
		),
		hotCache: cache.NewLoadingCache(
			func(key cache.Key) (value cache.Value, err error) {
				createTime := simple.Timestamp(time.Now().AddDate(0, 0, -3))
				value = repositories.ArticleRepository.Find(simple.DB(),
					simple.NewSqlCnd().Gt("create_time", createTime).Eq("status", constants.StatusOk).Desc("view_count").Limit(5))
				return
			},
			cache.WithMaximumSize(1),
			cache.WithRefreshAfterWrite(10*time.Minute),
		),
	}
}

func (c *articleCache) GetRecommendArticles() []model.Article {
	val, err := c.recommendCache.Get(articleRecommendCacheKey)
	if err != nil {
		return nil
	}
	if val != nil {
		return val.([]model.Article)
	}
	return nil
}

func (c *articleCache) InvalidateRecommend() {
	c.recommendCache.Invalidate(articleRecommendCacheKey)
}

func (c *articleCache) GetHotArticles() []model.Article {
	val, err := c.hotCache.Get(articleHotCacheKey)
	if err != nil {
		return nil
	}
	if val != nil {
		return val.([]model.Article)
	}
	return nil
}
