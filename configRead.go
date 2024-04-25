package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

func loadConfig(filename string) *Config {
	var config Config
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &config)

	if err != nil {
		panic(err)
	}

	return &config
}

type Config struct {
	Cars struct {
		Count           int `yaml:"count"`
		ArrivalTimeMin  int `yaml:"arrival_time_min"`
		ArrivalTimeMax  int `yaml:"arrival_time_max"`
		MainQueueLength int `yaml:"main_queue_length_max"`
	} `yaml:"cars"`
	Stations  map[FuelType]StationConfig `yaml:"stations"`
	Registers struct {
		Count          int `yaml:"count"`
		HandleTimeMin  int `yaml:"handle_time_min"`
		HandleTimeMax  int `yaml:"handle_time_max"`
		QueueLengthMax int `yaml:"queue_length_max"`
	} `yaml:"registers"`
}

type StationConfig struct {
	Count          int `yaml:"count"`
	ServeTimeMin   int `yaml:"serve_time_min"`
	ServeTimeMax   int `yaml:"serve_time_max"`
	QueueLengthMax int `yaml:"queue_length_max"`
}
