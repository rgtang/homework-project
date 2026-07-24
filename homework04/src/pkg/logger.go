package pkg

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {
	// 配置高可读性且性能极高的 JSON 格式 Encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        // 人类可读的时间格式：2026-07-23T01:05:51.000Z
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 控制台带颜色的日志级别

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // 终端彩色输出，生产环境可改为 NewJSONEncoder
		zapcore.AddSync(os.Stdout),
		zap.DebugLevel,
	)

	// 开启 Caller 记录（打印引发日志的文件名与行号）
	Logger = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(Logger) // 设置为全局 Zap 变量
}
