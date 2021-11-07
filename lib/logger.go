package lib

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

/*
type Logger struct {
}

var LOGGER Logger
*/

var Log = logrus.New()

func LogInit(confFilename string) {
	//conf := ConfigReader("../config/common.yaml")
	conf := ConfigReader(confFilename)
	fmt.Println(conf.ProcConf.Process.Data_path)
	fmt.Println(conf.CommonConf.Log.Filename)

	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.TraceLevel)

	/*
		Log.WithFields(logrus.Fields{
			"animal": "walrus",
			"size":   10,
		}).Info("A group of walrus emerges from the ocean")
	*/
}
