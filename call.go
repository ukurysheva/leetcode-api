package leetcode_api

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	url = "https://leetcode.com/graphql"
)

func callAPI(payload string) (*http.Response, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json; charset=UTF-8")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response status %s, code: %v", response.Status, response.StatusCode)
	}

	return response, nil
}
