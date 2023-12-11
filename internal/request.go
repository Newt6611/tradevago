package internal

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

// TODO: with header
func Get(ctx context.Context, requestUrl string, query url.Values) (b []byte, err error){
    req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
    if err != nil {
        return []byte{}, err
    }
    req.URL.RawQuery = query.Encode()

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return []byte{}, err
    }
    return io.ReadAll(res.Body)
}
