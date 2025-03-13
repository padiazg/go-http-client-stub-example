# http Client stub example
This is a simple example of how to use a stubbed HTTP client in Go for testing purposes.
Originaly found here: https://gist.github.com/cep21/1f2a5c61a2186db8d040b6ec20dd5db1, I just made it work with a simple example.

## Run the test
```shell
git clone https://github.com/padiazg/go-http-client-stub-example.git
cd go-http-client-stub-example
go test -timeout 30s -run .../.
```

## How to use
1. Create a new client using `NewClient()`
2. Set the transport of the client to a stubbed transport using `roundTripFn`
```go
    c = NewClient()
	c.Client.Transport = roundTripFn(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, r.URL.Path, "/v1/item/sword")
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"type":"sword"}`)),
		}, nil
	})
```
3. Call your function that uses the HTTP client
4. Assert the expected behavior in your test

## Acknowledgement
All credits goes to the maintainer of the repo mentioned above. This repo is for sharing the working example and for my own learning purposes.