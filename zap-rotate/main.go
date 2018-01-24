package main

import (
	"github.com/popwalker/studygo/zap-rotate/common/log"
)

func init() {
	log.InitLog()
}

func main() {
	// test log file
	for {
		log.LoggerSugar.Infof("Failed to fetch URL: %s", "https://www.baitu.com")
	}
}
