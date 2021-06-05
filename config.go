package main

import (
	"log"
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
		log.Fatal("sgits.yml config file cannot be found, place it in the working directory of sgits: ", err)
	}
	defer fp.Close()
	err = yaml.NewDecoder(fp).Decode(&config)
	if err != nil {
		log.Fatal("invalid yaml config: ", err)
		panic(err)
	}

	if config.Bin == "" {
		path, err := exec.Command("git", "--exec-path").Output()
		if err != nil {
			log.Fatal("error: cannot exec git, make sure your git is in your PATH: ", err)
		}
		config.Bin = string(path[0:len(path)-1]) + "/git-http-backend"
	}

	if _, err := os.Stat(config.Bin); err != nil {
		log.Fatal("error: cannot be found, make sure you have git installed: ", config.Bin)
	}
	log.Printf("found git-http-backend: %s", config.Bin)
}
