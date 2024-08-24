package main

import (
	"gopkg.in/yaml.v3"
	"os"
	"web-inspector-back/internal"
)

func main() {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	var config internal.Config

	if err = yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	app := internal.NewApplication(config)

	if err = app.Connect(); err != nil {
		panic(err)
	}

	defer app.CloseConnection()
}
