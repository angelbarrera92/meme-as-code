package main

import (
	"flag"
	"log"

	. "meme_as_code/src"
)

var (
	configPath = flag.String("config-file", "config.yaml", "The path to the config file.")
	override   = flag.Bool("override", false, "Override existing files.")
)

func main() {
	flag.Parse()

	config, err := GetConfigFromFile(*configPath)
	if err != nil {
		log.Fatal("Please, provide a valid integration.")
	}
	GetMemes(config)
}
