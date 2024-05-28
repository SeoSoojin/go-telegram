package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client interface {
	SendMessage(req RequestMessage) (Message, error)
	SendPhoto(req RequestPhoto) (Message, error)
}

type client struct {
	baseURL    *url.URL
	httpClient *http.Client
}

func NewClient(httpClient *http.Client, token string) (Client, error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	urlStr := fmt.Sprintf("https://api.telegram.org/bot%s/", token)
	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	return &client{
		httpClient: httpClient,
		baseURL:    baseURL,
	}, nil

}

func (c *client) SendMessage(message RequestMessage) (Message, error) {
	resp := new(Message)

	u := c.baseURL.ResolveReference(&url.URL{Path: "sendMessage"})

	b, err := json.Marshal(message)
	if err != nil {
		return *resp, err
	}

	httpReq, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(b))
	if err != nil {
		return *resp, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return *resp, err
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return *resp, fmt.Errorf("unexpected status code: %d", httpResp.StatusCode)
	}

	b, err = io.ReadAll(httpResp.Body)
	if err != nil {
		return *resp, err
	}

	err = json.Unmarshal(b, resp)
	if err != nil {
		return *resp, err
	}

	return *resp, nil
}

func (c *client) SendPhoto(photo RequestPhoto) (Message, error) {
	resp := new(Message)

	u := c.baseURL.ResolveReference(&url.URL{Path: "sendPhoto"})

	b, err := json.Marshal(photo)
	if err != nil {
		return *resp, err
	}

	httpReq, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(b))
	if err != nil {
		return *resp, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return *resp, err
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return *resp, fmt.Errorf("unexpected status code: %d", httpResp.StatusCode)
	}

	b, err = io.ReadAll(httpResp.Body)
	if err != nil {
		return *resp, err
	}

	err = json.Unmarshal(b, resp)
	if err != nil {
		return *resp, err
	}

	return *resp, nil
}
