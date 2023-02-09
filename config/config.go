package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Name   string `json:"name"`
	Branch string `json:"branch"`
}

var RunConf Config

func Init(runConf string) {
	jsonFile, err := os.Open(runConf)
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	defer jsonFile.Close()
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&RunConf)
	if err != nil {
		fmt.Println("Cannot get configuration from file")
		return
	}
}
