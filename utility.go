package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

func GetConfigValue(value, defaultValue string) string {
	file := "conf.ini"

	configFilePath, err := getConfigFilePath(file)
	if err != nil {
		fmt.Println("Error getting config file path:", err)
		return defaultValue
	}

	cfg, err := ini.Load(configFilePath)
	if err != nil {
		fmt.Println("Error loading config file:", err)
		return defaultValue
	}

	myValue := cfg.Section("").Key(value).String()
	if myValue != "" {
		return myValue
	}
	return defaultValue
}

func getConfigFilePath(fileName string) (string, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(workingDir, fileName), nil
}
