package lib

import (
	"fmt"
	"log"
	"os"
)

type LOGGER struct {
	logger *log.Logger
}

type Logger interface {
	debug() bool
	error() bool
	service() bool
}

func InitLogger() *LOGGER {
	conf := ConfigReader("../config/common.yaml")
	fmt.Println(conf.ProcConf.Process.Data_path)
	logger := LOGGER{}
	logger.logger = log.New(os.Stdout, "INFO : ", log.LstdFlags)
	return &logger
}

func (l *LOGGER) debug() bool {

}

func (l *LOGGER) error() bool {

}

func (l *LOGGER) service() bool {

}
