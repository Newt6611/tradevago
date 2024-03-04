package bito

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Newt6611/tradevago/pkg/api"
)

// https://github.com/bitoex/bitopro-offical-api-docs/blob/master/model.md#order-status-explanation
func getOrderStatus(s int) api.OrderStatus {
    switch (s) {
    case 0:
        return api.OrderStatusWait

    case 1:
        return api.OrderStatusDone
    case 2:
        return api.OrderStatusDone
    case 3:
        return api.OrderStatusDone

    case 4:
        return api.OrderStatusCancel
    default:
        return api.OrderStatusUnknow
    }
}

func newBitoAuthHeader(identity, apiKey, apiSecret string, method string, body map[string]interface{}) (http.Header, error) {
    var (
        payload string
        err     error
    )
    if method == "POST" {
        _, payload, err = getPostPayload(body)
        if err != nil {
            return nil, err
        }
    } else {
        payload = getNonPostPayload(identity, GetTimestamp())
    }
    sig := getSig(apiSecret, payload)

    header := http.Header{}
    header.Set("X-BITOPRO-APIKEY", apiKey)
    header.Set("X-BITOPRO-PAYLOAD", payload)
    header.Set("X-BITOPRO-SIGNATURE", sig)
    header.Set("X-BITOPRO-API", "golang")

    return header, nil
}

func getPostPayload(body map[string]interface{}) (string, string, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return "", "", err
	}
	return string(payload), base64.StdEncoding.EncodeToString(payload), nil
}

func getNonPostPayload(identity string, nonce int64) string {
	payload, _ := json.Marshal(map[string]interface{}{
		"identity": identity,
		"nonce":    nonce,
	})
	return base64.StdEncoding.EncodeToString(payload)
}

func getSig(secret, payload string) string {
	h := hmac.New(sha512.New384, []byte(secret))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}

func GetTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}
