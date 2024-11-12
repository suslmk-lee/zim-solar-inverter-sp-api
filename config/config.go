package config

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type AppConfigProperties map[string]interface{}

var ConfInfo AppConfigProperties

func init() {
	profile := os.Getenv("PROFILE")
	if profile == "" {
		profile = "dev" // 기본값 설정
	}

	if profile == "dev" {
		_, err := ReadConfigFile("config/config.json")
		if err != nil {
			log.Println("Failed to read config.json in dev mode:", err)
		}
	} else {
		ConfInfo = LoadConfigFromEnv()
	}
}

// ReadConfigFile reads configurations from a JSON file and stores them in ConfInfo
func ReadConfigFile(filename string) (AppConfigProperties, error) {
	ConfInfo = AppConfigProperties{}

	file, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening config.json:", err)
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&ConfInfo); err != nil {
		log.Println("Error decoding config.json:", err)
		return nil, err
	}

	return ConfInfo, nil
}

// LoadConfigFromEnv loads configurations from environment variables in production mode
func LoadConfigFromEnv() AppConfigProperties {
	conf := AppConfigProperties{}

	if origins := os.Getenv("ALLOWED_ORIGINS"); origins != "" {
		conf["AllowedOrigins"] = origins
	}

	return conf
}

// GetAllowedOrigins retrieves AllowedOrigins from the configuration
func GetAllowedOrigins() []string {
	if allowedOriginsInterface, exists := ConfInfo["AllowedOrigins"]; exists {
		switch v := allowedOriginsInterface.(type) {
		case string:
			return strings.Split(v, ",")
		case []interface{}:
			var origins []string
			for _, origin := range v {
				if strOrigin, ok := origin.(string); ok {
					origins = append(origins, strOrigin)
				}
			}
			return origins
		}
	}
	return nil
}
