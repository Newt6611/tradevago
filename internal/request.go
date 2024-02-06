package internal

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

func Get(ctx context.Context, requestUrl string, query url.Values, headers map[string]string) (b []byte, err error){
    req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
    if err != nil {
        return []byte{}, err
    }
    req.URL.RawQuery = query.Encode()
    for k, v := range headers {
        req.Header.Set(k, v)
    }

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return []byte{}, err
    }
    return io.ReadAll(res.Body)
}

func Post(ctx context.Context, requestUrl string, b []byte, headers map[string]string) ([]byte, error){
    req, err := http.NewRequest(http.MethodPost, requestUrl, bytes.NewReader(b))
    if err != nil {
        return []byte{}, err
    }
    for k, v := range headers {
        req.Header.Set(k, v)
    }

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return []byte{}, err
    }
    return io.ReadAll(res.Body)
}
