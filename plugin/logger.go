package plugin

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Log() *zap.Logger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./app.log",
		MaxSize:    20, // megabytes
		MaxBackups: 10,
		MaxAge:     7, // days
	})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)

	logger := zap.New(core)

	return logger

	// 範例
	// logger := plugin.Log()
	// defer logger.Sync()
	// logger.Info("userService.Create", zap.String("user.id", strconv.FormatUint(uint64(user.ID), 10)), zap.String("user.name", user.Name))
}
