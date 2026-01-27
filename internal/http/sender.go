package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Archiker-715/cache-server/internal/cache"
	"github.com/Archiker-715/cache-server/internal/constants"
)

func Send(port, url, method, body string) ([]byte, error) {
	var client = &http.Client{}
	resp, err := chooseRequestType(client, url, method, body)
	if err != nil {
		return nil, fmt.Errorf("send request error: %w", err)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response error: %w", err)
	}

	wasCached := cache.Cache(port, url, responseBody)
	if wasCached {
		resp.Header.Add("X-Cache", "HIT")
	} else {
		resp.Header.Add("X-Cache", "MISS")
	}

	return responseBody, nil
}

func sendWithBody(client *http.Client, url, method, body string) (*http.Response, error) {
	body = strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(body, "\n", ""), "\r", ""))

	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, fmt.Errorf("create req w body: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do req w body: %w", err)
	}

	return resp, nil
}

func sendWithoutBody(client *http.Client, url, method string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create req: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do req: %w", err)
	}

	return resp, nil
}

func chooseRequestType(client *http.Client, url, method, body string) (*http.Response, error) {
	haveBody := strings.EqualFold(body, "")
	if method == constants.GET && haveBody {
		return nil, fmt.Errorf("GET method cannot have body")
	}

	if haveBody {
		return sendWithBody(client, url, method, body)
	} else {
		return sendWithoutBody(client, url, method)
	}
}
