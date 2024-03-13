package log

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger
	initOnce      sync.Once
)

// Initialize the logger.
func InitLogger() {
	initOnce.Do(func() {
		configureLogger()
	})
}

// Load environment variables.

// Configure the logger.
func configureLogger() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logFilePath := os.Getenv("LOG_FILE")
	if logFilePath != "" {
		setupFileLogger(&config, logFilePath)
	} else {
		setupDefaultLogger(&config)
	}
}

// Setup logger with file output.
func setupFileLogger(config *zap.Config, logFilePath string) {
	logDir := "logs"
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		panic(fmt.Sprintf("Failed to create log directory: %v", err))
	}

	fullPath := filepath.Join(logDir, logFilePath)
	file, err := os.Create(fullPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to create log file: %v", err))
	}

	writeSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(file))
	encoder := zapcore.NewConsoleEncoder(config.EncoderConfig)
	core := zapcore.NewCore(encoder, writeSyncer, config.Level)
	logger = zap.New(core)
	sugaredLogger = logger.Sugar()
}

// Setup default logger.
func setupDefaultLogger(config *zap.Config) {
	var err error
	logger, err = config.Build()
	if err != nil {
		panic(fmt.Sprintf("Failed to initialise zap logger: %v", err))
	}
	sugaredLogger = logger.Sugar()
}

// Log functions.
func Info(msg string) {
	sugaredLogger.Info(msg)
}

func Infof(msg string, args ...interface{}) {
	sugaredLogger.Infof(msg, args...)
}

func Error(msg string) {
	sugaredLogger.Error(msg)
}

func Errorf(msg string, args ...interface{}) {
	sugaredLogger.Errorf(msg, args...)
}

func Warning(msg string) {
	sugaredLogger.Warn(msg)
}

func Warningf(msg string, args ...interface{}) {
	sugaredLogger.Warnf(msg, args...)
}

func Debug(msg string) {
	sugaredLogger.Debug(msg)
}

func Debugf(msg string, args ...interface{}) {
	sugaredLogger.Debugf(msg, args...)
}

func Panic(msg string) {
	sugaredLogger.Panic(msg)
}

func Panicf(msg string, args ...interface{}) {
	sugaredLogger.Panicf(msg, args...)
}
