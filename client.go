package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Client struct {
	Client *http.Client
}

type Data struct {
	Type string `json:"type"`
}

func NewClient() *Client {
	return &Client{Client: &http.Client{}}
}

func (c *Client) GetItem(ctx context.Context, field string) (*Data, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://localhost.com/v1/item/"+field, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data := &Data{}
	err = json.NewDecoder(res.Body).Decode(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
