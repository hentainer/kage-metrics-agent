
package main

import (
	"os"

	"gopkg.in/urfave/cli.v1"
)

import _ "github.com/joho/godotenv/autoload"

// Flag constants declared for CLI use.
const (
	FlagLog      = "log"
	FlagLogFile  = "log.file"
	FlagLogLevel = "log.level"

	FlagKafkaBrokers      = "kafka.brokers"
	FlagKafkaIgnoreTopics = "kafka.ignore-topics"
	FlagKafkaIgnoreGroups = "kafka.ignore-groups"

	FlagReporters = "reporters"

	FlagInflux       = "influx"
	FlagInfluxMetric = "influx.metric"
	FlagInfluxPolicy = "influx.policy"
	FlagInfluxTags   = "influx.tags"

	FlagServer = "server"
	FlagPort   = "port"
)

// Version is the compiled application version.
var Version = "¯\\_(ツ)_/¯"

var commonFlags = []cli.Flag{
	cli.StringFlag{
		Name:   FlagLog,
		Value:  "stdout",
		Usage:  "The type of log to use (options: \"stdout\", \"file\")",
		EnvVar: "KAGE_LOG",
	},
	cli.StringFlag{
		Name:   FlagLogFile,
		Usage:  "The path to the file",
		EnvVar: "KAGE_LOG_FILE",
	},
	cli.StringFlag{
		Name:   FlagLogLevel,
		Value:  "info",
		Usage:  "Specify the log level (options: \"debug\", \"info\", \"warn\", \"error\")",
		EnvVar: "KAGE_LOG_LEVEL",
	},
}
