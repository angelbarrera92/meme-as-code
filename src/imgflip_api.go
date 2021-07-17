package meme_as_code

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

const ImgApiUrl = "https://api.imgflip.com/caption_image"

func GetMemes(config *Config) error {
	//create output directory
	err := os.MkdirAll(config.OutputDir, 0775)
	if err != nil {
		log.Fatal(err)
	}
	for _, meme := range config.Memes {
		imageUrl, err := imgFlipApiPost(ImgApiUrl, meme, config.Username, config.Password)
		if err != nil {
			log.Print(err)
			break
		}
		err = downloadFile(imageUrl,
			config.OutputDir+"/"+meme.Filename,
			config.Overrive)
		if err != nil {
			return err
		}
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
	apiRawResp, err := http.PostForm(apiUrl, data)
	if err != nil {
		return "", err
	}

	imageUrl, err := GetUrl(apiRawResp.Body)
	if err != nil {
		return "", err
	}
	return imageUrl, nil
}

func GetUrl(apiRawResp io.Reader) (string, error) {
	var response map[string]interface{}

	json.NewDecoder(apiRawResp).Decode(&response)
	if success, ok := response["success"].(bool); ok {
		if !success {
			if error_message, ok := response["error_message"].(string); ok {
				return "", errors.New(error_message)
			} else {
				return "", errors.New("Unexpected response format for 'error_message' field")
			}
		} else {
			if data, ok := response["data"].(map[string]interface{}); ok {
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
	// Checks if the file exists
	if _, err := os.Stat(path); err == nil {
		if !override {
			return errors.New("File already exists")
		}
	}
	// Download the file
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}
