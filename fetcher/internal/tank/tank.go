package tank

import (
	"context"

	"freelancer.com/pump-over-images/fetcher/pkg/fetch"
)

type CrawlProperties struct {
	Url     string
	Method  string
	Pointer int64
	Payload map[string]interface{}
}

type Tank struct {
	f fetch.Fetch
}

type T interface {
	Crawl(context.Context, int64) (map[string]interface{}, error)
}

func NewTank(ctx context.Context, f fetch.Fetch) *Tank {
	return &Tank{
		f: f,
	}
}

func (t *Tank) Crawl(ctx context.Context, properties CrawlProperties) (map[string]interface{}, error) {
	// 1. pull image URI from API third-party
	// 2. filter URI to avoid duplicate via hash encode
	// 3. redis stream
	return nil, nil
}
