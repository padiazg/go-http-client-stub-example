package main

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// taken from :
// https://gist.github.com/cep21/1f2a5c61a2186db8d040b6ec20dd5db1

type roundTripFn func(r *http.Request) (*http.Response, error)

func (s roundTripFn) RoundTrip(r *http.Request) (*http.Response, error) {
	return s(r)
}

func TestSword(t *testing.T) {
	var (
		c           = NewClient()
		ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	)
	defer cancel()

	c.Client.Transport = roundTripFn(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, r.URL.Path, "/v1/item/sword")
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"type":"sword"}`)),
		}, nil
	})

	item, err := c.GetItem(ctx, "sword")
	assert.Nil(t, err)
	assert.Equal(t, item.Type, "sword")
}
