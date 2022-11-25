package service

import (
	"io"
	"net/http"
)

func GetJson(url string) string {
	var client http.Client
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := client.Do(req)
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}
