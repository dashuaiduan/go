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

func Test2(t *testing.T) {
	//	记录层级关系，子记录器 ，预设字段
	//	logger,_ := zap.NewDevelopment()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("tracked some metrics1111",
		zap.Namespace("metrics"),
		zap.Int("counter", 1),
	)
	logger.Info("tracked some metrics222")
	//子记录器 ，预设字段
	logger2 := logger.With(
		zap.Int("counter", 1),
	)
	logger2.Info("tracked some metrics")
}
