package logger

import (
	"strings"

	"github.com/tr2vil/its_manager/lib/config"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log = logrus.New()

func LogInit(confFilename string) {
	var log_level logrus.Level
	var log_filename string
	var lum *lumberjack.Logger

	conf := config.ConfigReader(confFilename)

	log_filename = conf.CommonConf.Log.Filename

	// Set Log Rotate
	lum = &lumberjack.Logger{
		Filename:   log_filename,
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     14,
		Compress:   true,
	}

	// Set Log Level
	switch strings.ToLower(conf.CommonConf.Log.Level) {
	case "panic":
		log_level = logrus.PanicLevel
	case "fatal":
		log_level = logrus.FatalLevel
	case "error":
		log_level = logrus.ErrorLevel
	case "warn":
		log_level = logrus.WarnLevel
	case "info":
		log_level = logrus.InfoLevel
	case "debug":
		log_level = logrus.DebugLevel
	case "trace":
		log_level = logrus.TraceLevel
	default:
		log_level = logrus.ErrorLevel
	}

	//Log.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339})
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceQuote:    true,
	})
	Log.SetOutput(lum)
	Log.SetLevel(log_level)
	Log.SetReportCaller(true)

	/*
		Log.WithFields(logrus.Fields{
			"animal": "walrus",
			"size":   10,
		}).Info("A group of walrus emerges from the ocean")
	*/
}
