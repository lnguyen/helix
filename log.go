package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/logutils"
)

const EnvLog = "HELIX_LOG"

func setLog() {
	log.SetOutput(ioutil.Discard)
	envLog := strings.ToUpper(os.Getenv(EnvLog))
	if envLog != "" {
		filter := &logutils.LevelFilter{
			Levels:   []logutils.LogLevel{"martini", "DEBUG", "WARN", "ERROR", "INFO"},
			MinLevel: logutils.LogLevel(envLog),
			Writer:   os.Stderr,
		}
		log.SetOutput(filter)
	}
}
