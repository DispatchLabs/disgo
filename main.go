package main

import (
	"os"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"encoding/json"
	"github.com/dispatchlabs/disgo/configurations"
	"github.com/dispatchlabs/disgo/server"
)

func main() {

	// Setup log.
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	// Read configuration JSON file.
	filePath := "." + string(os.PathSeparator) + "configurations" + string(os.PathSeparator) + "disgo_config.json"
	file, error := ioutil.ReadFile(filePath)
	if error != nil {
		log.Error("unable to load " + filePath)
		os.Exit(1)
	}
	json.Unmarshal(file, &configurations.Configuration)

	server := &server.Server{}
	server.Start()
}