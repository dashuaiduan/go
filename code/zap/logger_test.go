package zap

import (
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestName5544ff(t *testing.T) {
	//正式项目 这两个变量 应该是全局变量 共享使用
	logger, sugar := New(zap.DebugLevel, "./", "test")
	sugar.Errorf("test errof %s", "ggggg")
	sugar.Debugf("test debuf %s", "ggggg")
	sugar.Warnf("test Warnf %s", "ggggg")
	logger.Info("failed to fetch URL",
		zap.String("url", "https://fffffff"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second*10),
	)
}
