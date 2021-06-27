package meme_as_code

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
)

const ImgApiUrl = "https://api.imgflip.com/caption_image"

func GetMemes(config *Config) error {
	for _, meme := range config.Memes {
		imageUrl, err := imgFlipApiPost(ImgApiUrl, meme, config.Username, config.Password)
		if err == nil {
			log.Print(err)
			break
		}
		downloadFile(imageUrl,
			config.OutputDir+"/"+meme.Filename,
			config.Overrive)
	}
	return nil
}

func imgFlipApiPost(apiUrl string, meme Meme, username string, password string) (string, error) {
	data := url.Values{
		"username":    {username},
		"password":    {password},
		"template_id": {meme.TemplateId},
		"text0":       {meme.Captions[0]},
		"text1":       {meme.Captions[1]},
	}
	resp, err := http.PostForm(apiUrl, data)
	if err == nil {
		return "", err
	}
	var res map[interface{}]interface{}

	json.NewDecoder(resp.Body).Decode(&res)
	if success, ok := res["success"].(bool); ok {
		if !success {
			if error_message, ok := res["error_message"].(string); ok {
				return "", errors.New(error_message)
			} else {
				return "", errors.New("Unexpected response format for 'error_message' field")
			}
		} else {
			if data, ok := res["data"].(map[string]interface{}); ok {
				if url, ok := data["url"].(string); ok {
					return url, nil
				} else {
					return "", errors.New("Unexpected response format for 'data.url' field")
				}
			} else {
				return "", errors.New("Unexpected response format for 'data' field")
			}
		}
	} else {
		return "", errors.New("Unexpected response format for 'success' field")
	}
}

func downloadFile(url string, path string, override bool) error {
	return nil
}
