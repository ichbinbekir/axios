package axios

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type AxiosRequestConfig struct {
	Method  string
	Url     string
	Headers map[string]string
	Data    any
}

type AxiosResponse struct {
	Status     int
	StatusText string
	Headers    map[string]string
	Data       any
	//Config InternelAxiosRequestConfig
	//Request any
}

func Get(url string, config AxiosRequestConfig) (*AxiosResponse, error) {
	config.Method = "get"
	config.Url = url
	return Request(config)
}

func Post(url string, data any, config AxiosRequestConfig) (*AxiosResponse, error) {
	config.Method = "post"
	config.Url = url
	config.Data = data
	return Request(config)
}

func Request(config AxiosRequestConfig) (*AxiosResponse, error) {
	reqBuffer, err := json.Marshal(config.Data)
	if err != nil {
		return nil, err
	}

	requestHttp, err := http.NewRequest(config.Method, config.Url, bytes.NewReader(reqBuffer))
	if err != nil {
		return nil, err
	}

	for key, value := range config.Headers {
		requestHttp.Header.Add(key, value)
	}

	responseHttp, err := http.DefaultClient.Do(requestHttp)
	if err != nil {
		return nil, err
	}

	resBuffer, err := io.ReadAll(responseHttp.Body)
	if err != nil {
		return nil, err
	}

	response := new(AxiosResponse)
	response.Data = string(resBuffer)
	response.Status = responseHttp.StatusCode
	response.StatusText = responseHttp.Status

	response.Headers = make(map[string]string)
	for key, value := range responseHttp.Header {
		response.Headers[key] = value[0]
	}

	return response, nil
}
