package main

import (
	"flag"
	"log"
	"os"

	. "meme_as_code/src"
)

var (
	configPath = flag.String("config-file", "config.yaml", "The path to the config file.")
	outputDir  = flag.String("output-dir", "output", "The path where to download memes.")
	override   = flag.Bool("override", false, "Override existing files.")
)

func main() {
	flag.Parse()

	config, err := GetConfigFromFile(*configPath)
	if err != nil {
		log.Fatal("Please, provide a valid configuration.")
	}
	// get user and password from env
	config.Username = os.Getenv("USER")
	config.Password = os.Getenv("PASS")
	if config.Username == "" || config.Password == "" {
		log.Fatal("Please, provide a valid username and password from environment variables USER and PASS")
	}
	config.OutputDir = *outputDir
	config.Override = *override
	err = GetMemes(config)
	if err != nil {
		log.Fatal(err)
	}
}
