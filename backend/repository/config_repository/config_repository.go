package config_repository

import (
	"os"
	"encoding/json"
	"io/ioutil"
	"log"
	"image/color"
)

const (
	DISPLAY_CONFIG_FILE = "config/display_config.json"
)

type WordclockConfig struct {
	DisplayConfig Config
}

type Config struct {
	Color color.Color
	// Whether the flow should be alternated
	// This is caused by soldering the strips at the ends
	// Row1 :    1 2 3 4 5 6 7 __
	//                           |<= Soldering
	// Row 2:  __. . . . . 9 8 __|<=
	//        |
	// Row 3: |__. . . . . . .
	AlternateFlow bool
	// Whether to alternate the flow on even rows.
	// If false, the flow will be alternated on uneven rows
	// The lines of the display are counted with a starting index of 1
	AlternateEven bool
}

var config WordclockConfig
var initialized bool

func GetConfig() WordclockConfig {
	if !initialized {
		if _, err := os.Stat(DISPLAY_CONFIG_FILE); err == nil {
			return readConfigFile()
		} else {
			log.Printf("Config file %s not found. Use default config.", DISPLAY_CONFIG_FILE)
			return getDefaultConfig()
		}
	}
	return config
}

func SetConfig(c WordclockConfig) {
	config = c
	go writeConfigFile(c)
}

func writeConfigFile(c WordclockConfig) {
	jsonData, err := json.Marshal(c)
	if err != nil {
		log.Printf("Error when marshalling config: %v", err)
		return
	}

	if err = ioutil.WriteFile(DISPLAY_CONFIG_FILE, jsonData, 0666); err != nil {
		log.Printf("Error when updating config: %v", err)
	}
}

func readConfigFile() WordclockConfig {
	bytes, err := ioutil.ReadFile(DISPLAY_CONFIG_FILE)
	if err != nil {
		log.Printf("Error when reading from file %s: %v", DISPLAY_CONFIG_FILE, err)
		log.Print("Using default config")
		return getDefaultConfig()
	}
	config := WordclockConfig{}
	json.Unmarshal(bytes, &config)
	return config
}

func getDefaultConfig() WordclockConfig {
	return WordclockConfig{
		DisplayConfig: Config{
			Color:         color.White,
			AlternateFlow: true,
			AlternateEven: true,
		},
	}
}
