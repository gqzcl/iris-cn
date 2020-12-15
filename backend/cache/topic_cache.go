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
	topicRecommendCacheKey = "recommend_topics_cache"
)

var TopicCache = newTopicCache()

type topicCache struct {
	recommendCache cache.LoadingCache
}

func newTopicCache() *topicCache {
	return &topicCache{
		recommendCache: cache.NewLoadingCache(
			func(key cache.Key) (value cache.Value, e error) {
				value = repositories.TopicRepository.Find(simple.DB(),
					simple.NewSqlCnd().Eq("status", constants.StatusOk).Desc("id").Limit(50))
				return
			},
			cache.WithMaximumSize(10),
			cache.WithRefreshAfterWrite(30*time.Minute),
		),
	}
}

func (c *topicCache) GetRecommendTopics() []model.Topic {
	val, err := c.recommendCache.Get(topicRecommendCacheKey)
	if err != nil {
		return nil
	}
	if val != nil {
		return val.([]model.Topic)
	}
	return nil
}

func (c *topicCache) InvalidateRecommend() {
	c.recommendCache.Invalidate(topicRecommendCacheKey)
}
