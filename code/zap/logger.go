package zap

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// 实现日志分割与轮转 的zap

// 返回zap的两种形态

func New(level zapcore.Level, rootPath, serviceName string) (logger *zap.Logger, loggerSugar *zap.SugaredLogger) {
	fmt.Println("rootPath----", rootPath)
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	//encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	})

	// 实现两个判断日志等级的interface,可以添加修改条件
	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.DebugLevel
	})

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel && lvl < zapcore.ErrorLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	// 获取 info、error日志文件的io.Writer 抽象 getWriter() 在下方实现
	infoWriter := getWriter(filepath.Join(rootPath, "logs", "info.log"))
	errorWriter := getWriter(filepath.Join(rootPath, "logs", "error.log"))
	debugWriter := getWriter(filepath.Join(rootPath, "logs", "debug.log"))

	Alevel := zap.NewAtomicLevel()
	Alevel.SetLevel(level)
	// 最后创建具体的Logger
	var core zapcore.Core
	if rootPath != "no" {
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(debugWriter), debugLevel),
			//控制台输出日志
			zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), Alevel),
		)
	} else { //  ==no  不写文件 只打印到控制台
		core = zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), Alevel)
	}

	// 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 配置AddCallerSkip才能打印的原始行号
	logger = zap.New(core, zap.AddCaller(), zap.Development(), zap.Fields(zap.String("serviceName", serviceName)), zap.AddCallerSkip(1), zap.Hooks(hookDemo))
	loggerSugar = logger.Sugar()
	return
}

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		strings.Replace(filename, ".log", "", -1) + "-%Y%m%d%H.log", // 没有使用go风格反人类的format格式
		//rotatelogs.WithLinkName(filename),
		//rotatelogs.WithMaxAge(time.Hour*24*7),
		//rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		panic(err)
	}

	return hook
}

// hook 测试
func hookDemo(entry zapcore.Entry) error {
	fmt.Println("hook:----------", entry.Level, entry.Message)
	return nil
}
