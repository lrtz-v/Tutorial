package main

import (
	"fmt"
	"os"
	"path/filepath"

	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *Log

type Log struct {
	logger    *zap.Logger
	logFile   string
	logWFFile string
}

func (log *Log) Info(message string) {
	log.logger.Info(message)
}

func (log *Log) InfoF(format string, a ...interface{}) {
	log.logger.Info(fmt.Sprintf(format, a...))
}

func (log *Log) Debug(message string) {
	log.logger.Debug(message)
}

func (log *Log) DebugF(format string, a ...interface{}) {
	log.logger.Debug(fmt.Sprintf(format, a...))
}

func (log *Log) Warn(message string) {
	log.logger.Warn(message)
}

func (log *Log) WarnF(format string, a ...interface{}) {
	log.logger.Warn(fmt.Sprintf(format, a...))
}

func (log *Log) Error(message string) {
	log.logger.Error(message)
}

func (log *Log) ErrorF(format string, a ...interface{}) {
	log.logger.Error(fmt.Sprintf(format, a...))
}

func logDirCreate(logDir string) error {
	return DirGenrate(logDir)
}

func generateLogFiles(logFilePath string) string {
	dir := strings.TrimSuffix(logFilePath, "/")
	logFile := filepath.Join(dir, LogFileName)
	if _, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE, 0666); err != nil {
		panic(err)
	}
	return logFile
}

func hook(entry zapcore.Entry, samplingDecision zapcore.SamplingDecision) {
	fmt.Printf("[%s] %s\n", entry.Level, entry.Message)
}

func InitLogger(logFilePath string) {

	if err := logDirCreate(logFilePath); err != nil {
		fmt.Printf("[ERROR] Create log dir failed, err: %s\n", err.Error())
	}

	logFile := generateLogFiles(logFilePath)

	conf := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: false,
		Encoding:    "json",
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
			Hook:       hook,
		},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "Date",
			LevelKey:       "Level",
			NameKey:        "Key",
			CallerKey:      "Line",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "Message",
			StacktraceKey:  "Stack",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths: []string{logFile},
	}
	zapLogger, err := conf.Build()
	if err != nil {
		panic(err)
	}

	Logger = &Log{logger: zapLogger, logFile: logFile, logWFFile: "logWFFile"}
}
