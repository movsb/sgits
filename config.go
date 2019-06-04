package main

import (
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Listen   string `yaml:"listen"`
	Root     string `yaml:"root"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Bin      string `yaml:"bin"`
}

var config Config

func mustParseConfig(file string) {
	fp, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	err = yaml.NewDecoder(fp).Decode(&config)
	if err != nil {
		panic(err)
	}

	if config.Bin == "" {
		path, err := exec.Command("git", "--exec-path").Output()
		if err != nil {
			panic(err)
		}
		config.Bin = string(path[0:len(path)-1]) + "/git-http-backend"
	}
}
