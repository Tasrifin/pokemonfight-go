package helpers

import (
	"crypto/tls"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Tasrifin/pokemonfight-go/constants"
)

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

	return result, nil
}

func GetIDByUrl(url string) (result int, err error) {
	remover := constants.BASE_API_URL + "/" + constants.POKEMON_URI + "/"
	result1 := strings.ReplaceAll(url, remover, "")
	result2 := strings.ReplaceAll(result1, "/", "")

	result, err = strconv.Atoi(result2)
	return
}

func CheckDuplicateID(pokemonIDs []int) bool {
	temp := []int{}
	for _, v := range pokemonIDs {
		for _, w := range temp {
			if v == w {
				return true
			}
		}
		temp = append(temp, v)
	}

	return false
}
