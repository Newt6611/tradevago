package max

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Newt6611/tradevago/pkg/api"
)

func generateNonce() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func generateSignature(secret, nonce string) string {
	hmacHash := hmac.New(sha256.New, []byte(secret))
	hmacHash.Write([]byte(nonce))
	return hex.EncodeToString(hmacHash.Sum(nil))
}

func getOrderStatus(state string) api.OrderStatus {
    switch state {
    case "wait":
        return api.OrderStatusWait
    case "done":
        return api.OrderStatusDone
    case "cancel":
        return api.OrderStatusCancel
    }

    return api.OrderStatusUnknow
}
func createPayload(params map[string]interface{}, path string) (string, error) {
	paramsToBeSigned := make(map[string]interface{})
	for key, value := range params {
		paramsToBeSigned[key] = value
	}
	paramsToBeSigned["path"] = path

	payloadBytes, err := json.Marshal(paramsToBeSigned)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(payloadBytes), nil
}

func createSignatureWithPayload(payload string, secretKey string) string {
	key := []byte(secretKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(payload))
	return fmt.Sprintf("%x", h.Sum(nil))
}
