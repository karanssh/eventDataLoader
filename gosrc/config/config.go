package config

import (
	"encoding/json"
	"log"
	"os"
)

type (
	configType struct {
		DBUserName           string
		DBPassword           string
		DBHost               string
		DBPort               string
		DBProtocol           string
		DBName               string
		TableName            string
		IngestFolderLocation string
		PersistData          bool
		ClearDataBeforeRun   bool
	}
)

var (

	// Config contains the config items
	Config = configType{
		DBUserName:           "root",
		DBPassword:           "karanTest",
		DBHost:               "127.0.0.1",
		DBPort:               "3306",
		DBProtocol:           "tcp",
		DBName:               "ingestDB",
		TableName:            "ingestEvents",
		IngestFolderLocation: "/home/karan/projects/netcracker/generateEventData/outputData",
		PersistData:          true,
		ClearDataBeforeRun:   true,
	}
)

func loadJSON(jsonFile string) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	// log.Println("path is " + path)
	log.Println("path of config file is " + path + jsonFile)
	f, err := os.Open(path + jsonFile)
	if nil != err {
		log.Printf("Error loading app configuration: err=%v", err)
		return
	}

	dec := json.NewDecoder(f)

	dec.Decode(&Config)
}

func init() {
	loadJSON("/confVolume/config.json")
}
