package meme_as_code

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Meme struct {
	Filename   string   `yaml:"filename"`
	TemplateId string   `yaml:"template_id"`
	Captions   []string `yaml:"captions"`
}

type Config struct {
	OutputDir string `default:""`
	Username  string `default:""`
	Password  string `default:""`
	Override  bool   `default:false`
	Memes     []Meme `yaml:"memes"`
}

// GetConfigFromFile Returns a Config from a file
func GetConfigFromFile(path string) (*Config, error) {
	var configParsed Config

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = yaml.NewDecoder(file).Decode(&configParsed)
	if err != nil {
		return nil, err
	}
	return &configParsed, nil
}
