package lunchmoney

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Client ...
type Client struct {
	accessToken string
	baseURL     string
	http        *http.Client
}

// NewClient ...
func NewClient(accessToken string) (*Client, error) {
	return &Client{
		accessToken: accessToken,
		baseURL:     "https://dev.lunchmoney.app/v1",
		http:        &http.Client{Timeout: time.Minute * 3},
	}, nil
}

// Call ...
func (client *Client) Call(method, endpoint string, body []byte, v interface{}) error {
	request, err := client.newRequest(method, endpoint, bytes.NewReader(body), v)
	if err != nil {
		return err
	}

	return client.do(request, v)
}

func (client *Client) newRequest(method, endpoint string, body io.Reader, v interface{}) (*http.Request, error) {
	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}

	req, err := http.NewRequest(method, string(client.baseURL)+endpoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.accessToken))

	return req, nil
}

func (client *Client) do(req *http.Request, v interface{}) error {
	res, err := client.http.Do(req)

	if err != nil {
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode == 200 {
		return json.NewDecoder(res.Body).Decode(v)
	}

	return nil
}
