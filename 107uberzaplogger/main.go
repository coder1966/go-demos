package main

/*
	go get go.uber.org/zap
	go get -u github.com/natefinch/lumberjack

	/ DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel Level = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
*/

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var l *zap.SugaredLogger

func main() {
	initLogger(zap.DebugLevel)
	defer func() { _ = l.Sync() }() // 本来是 defer l.Sync()，不这么写 lint 不过
	l.Error("============== 开始 ===========")
	sub()
	l.Error("============== 换级别，再实验 ===========")
	// 换级别，再实验
	initLogger(zap.ErrorLevel)
	sub()
	l.Error("============== 结束 ===========")
	l.Fatal("l.Fatal ", "测试")
	fmt.Println("main 结束")
}

func initLogger(level zapcore.LevelEnabler) {
	encoder := getEncoder()
	writeSyncers := getLogWriter()
	// Encoder:编码器(如何写入日志)。WriterSyncer ：指定日志将写到哪里去。Level：哪种级别的日志将被写入。DebugLevel
	core := zapcore.NewCore(encoder, writeSyncers, level)
	logger := zap.New(core, zap.AddCaller()) // AddCaller()为显示文件名和行号
	// logger := zap.New(core) // AddCaller()为显示文件名和行号
	l = logger.Sugar()
}

// Encoder:编码器(如何写入日志)。
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // 修改时间编码器
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 大写记录日志级别
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 在zap中加入Lumberjack支持，切割日志文件。
func getLogWriter() zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "./test.log", // 日志文件的位置
		MaxSize:    1,            // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 5,            // 保留旧文件的最大个数
		MaxAge:     30,           // 保留旧文件的最大天数
		Compress:   true,         // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberjackLogger)
}

func sub() {
	l.Debug("l.Debug ", "测试")
	l.Debugln("l.Debugln ", "测试")
	l.Info("l.Info ", "测试")
	l.Infoln("l.Infoln ", "测试")
	l.Warn("l.Warn ", "测试")
	l.Error("l.Error ", "测试")
	sub0()
	// l.Panic("l.Panic ", "测试")
	// l.Fatal("l.Fatal ", "测试")
}

func sub0() {
	defer func() {
		if errPanic := recover(); errPanic != nil {
			fmt.Println("errPanic", errPanic)
		} else {
			fmt.Println("else errPanic", errPanic)
		}
	}()
	sub00()
}

func sub00() {
	defer func() {
		if errPanic := recover(); errPanic != nil {
			fmt.Println("errPanic", errPanic)
			l.Panic("sub00() l.Panic ", "测试")
		} else {
			fmt.Println("else errPanic", errPanic)
			l.Info("sub00() not l.Panic ", "测试")
		}
	}()
	_ = sub01()
}

func sub01() string {
	for {
		err := sub02()
		if err == nil {
			// Need get data to find backend errors.
			err = nil
			if err == nil {
				return "正常返回"
			}
			// Sleep, retry.
		}
		time.Sleep(time.Second * 2)
	}
}

func sub02() error { a := 0; _ = a; panic("人为的恐慌") } // 还可以这么写代码，全在一行里。还可以这么写代码，全在一行里。还可以这么写代码，全在一行里。还可以这么写代码，全在一行里。还可以这么写代码，全在一行里。还可以这么写代码，全在一行里。还可以这么写代码，全在一行里。
