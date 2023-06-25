package helpers

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func GetENV(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func ReqHTTP(method string, url string) (result *http.Response, err error) {
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: time.Duration(10) * time.Second,
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	result, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		err := fmt.Errorf("request status error : %v", result.StatusCode)
		return nil, err
	}

	return result, nil
}
