package utils

import (
	"errors"
	"io"

	"scon/config"
)

type Response struct {
	Data string
}

func HttpGet(api string) (string, error) {
	if len(api) == 0 {
		return "", errors.New("api is empty")
	}

	respRaw, err := config.HttpClient.Get(api)
	defer respRaw.Body.Close()
	if err != nil {
		return "", err
	}

	if respRaw.StatusCode != 200 {
		return "", errors.New("server error")
	}

	bodyBytes, err := io.ReadAll(respRaw.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}
