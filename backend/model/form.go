package model

// CreateCommentForm 发表评论
type CreateCommentForm struct {
	EntityType  string `form:"entityType"`
	EntityID    int64  `form:"entityId"`
	Content     string `form:"content"`
	QuoteID     int64  `form:"quoteId"`
	ContentType string `form:"contentType"`
}
