package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
)

var (
	Logger *zap.Logger
	LoggerSugar *zap.SugaredLogger
)

type WriteSyncer struct{
	io.Writer
}

func (ws WriteSyncer)Sync()error {
	return nil
}

func InitLog(){
	var cfg zap.Config
	var runMode = "debug"
	var logName = "logs/app.log"

	if runMode == "release"{
		cfg = zap.NewProductionConfig()
		cfg.DisableCaller = true
	}else {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.LevelKey = "level"
		cfg.EncoderConfig.NameKey = "name"
		cfg.EncoderConfig.MessageKey = "msg"
		cfg.EncoderConfig.CallerKey = "caller"
		cfg.EncoderConfig.StacktraceKey = "stacktrace"
	}

	cfg.Encoding = "json"
	cfg.EncoderConfig.TimeKey = "ts"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.OutputPaths =[]string{logName}
	//cfg.ErrorOutputPaths = []string{"logs/error.log"}

	sw := getWriteSyncer(logName)

	l, err := cfg.Build(SetOutput(sw, cfg))
	if err != nil {
		panic(err)
	}
	defer l.Sync()

	Logger = l
	sugar := Logger.Sugar()
	LoggerSugar = sugar.Desugar().WithOptions(SetOutput(sw, cfg)).Sugar()
}

// SetOutput replaces existing Core with new, that writes to passed WriteSyncer.
func SetOutput(ws zapcore.WriteSyncer, conf zap.Config) zap.Option{
	var enc zapcore.Encoder
	switch conf.Encoding {
	case "json":
		enc = zapcore.NewJSONEncoder(conf.EncoderConfig)
	case "console":
		enc = zapcore.NewConsoleEncoder(conf.EncoderConfig)
	default:
		panic("unknown encoding")
	}

	return zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewCore(enc, ws, conf.Level)
	})
}

// 设置利用lumberjack实现zapcore.WriteSyncer
func getWriteSyncer(logName string) zapcore.WriteSyncer{
	var ioWriter  = &lumberjack.Logger{
		Filename:   logName,
		MaxSize:    20, // MB
		MaxBackups: 3, // number of backups
		MaxAge:     28, //days
		LocalTime:  true,
		Compress:   false, // disabled by default
	}
	var sw = WriteSyncer{
		ioWriter,
	}
	return sw
}

