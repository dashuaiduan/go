package zap

import (
	"go.uber.org/zap"
	"testing"
	"time"
)

func Test1(t *testing.T) {

	logger := zap.NewExample()
	defer logger.Sync()

	url := "http://example.org/api"
	logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second*10),
	)
	//这种方式简单直观 但是 性能下降50%
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
