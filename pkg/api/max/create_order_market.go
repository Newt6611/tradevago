package max

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/Newt6611/tradevago/internal"
	"github.com/Newt6611/tradevago/pkg/api"
)

type createMarketOrderRequest struct {
	Side        string `json:"side,omitempty"`
	OrdType     string `json:"ord_type,omitempty"`
	Market      string `json:"market,omitempty"`
	Volume      string `json:"volume,omitempty"`
}
type createOrderResponse struct {
	ID                int64     `json:"id"`
	ClientOID         *string   `json:"client_oid"`
	Side              string    `json:"side"`
	OrderType         string    `json:"ord_type"`
	Price             *float64  `json:"price"`
	StopPrice         *float64  `json:"stop_price"`
	State             string    `json:"state"`
	Market            string    `json:"market"`
	CreatedAtInMs     int64     `json:"created_at_in_ms"`
	UpdatedAtInMs     int64     `json:"updated_at_in_ms"`
	Volume            string    `json:"volume"`
	RemainingVolume   string    `json:"remaining_volume"`
	ExecutedVolume    string    `json:"executed_volume"`
	TradesCount       int       `json:"trades_count"`
	GroupID           *string   `json:"group_id"`
}

func (this *Max) CreateOrderMarket(ctx context.Context, side api.Side, pair string, baseAmount float64, quoteAmount float64) (api.Order, error) {
    path := "/api/v2/orders"
    var maxside string
    if side == api.SELL {
        maxside = "sell"
    } else if side == api.BUY {
        maxside = "buy"
    }
    if baseAmount == 0 {
        return api.Order{}, errors.New("[MAX] Create order market should not be 0")
    }

    //-------------------------header---------------------------------------//
    headers := make(map[string]string)
    params := map[string]interface{}{
        "nonce": generateNonce(),
    }
    payload, err := createPayload(params, path)
    if err != nil {
        return api.Order{}, err
    }
    headers["Content-Type"] = "application/json"
    headers["X-MAX-ACCESSKEY"] = this.apiKey
    headers["X-MAX-PAYLOAD"] = payload
    headers["X-MAX-SIGNATURE"] = createSignatureWithPayload(payload, this.apiSecret)
    //---------------------------------------------------------------------//
    createOrder := createMarketOrderRequest {
        Side: maxside,
        OrdType: "market",
        Volume: strconv.FormatFloat(baseAmount, 'f', -1, 64),
        Market: pair,
    }
    reqb, err := json.Marshal(createOrder)
    if err != nil {
        return api.Order{}, err
    }
    res, err := internal.Post(ctx, MAX_API_ENDPOINT + path, reqb, headers)

    var errResponse apiErrorResponse
    err = json.Unmarshal(res, &errResponse)
	if err == nil && errResponse.Error.Code == 2018 {
		return api.Order{}, api.ErrorBalanceNotEnougth
	}


    var resOrder createOrderResponse
    err = json.Unmarshal(res, &resOrder)
	if err != nil {
		return api.Order{}, err
	}

    returnOrder := mapOrder(resOrder)
    if returnOrder.OrderStatus == api.OrderStatusUnknow && returnOrder.Id == "0" {
        return api.Order{}, errors.New(string(res))
    }
    return mapOrder(resOrder), nil
}

func mapOrder(order createOrderResponse) api.Order {
    o := api.Order {
        Id: fmt.Sprintf("%d", order.ID),
        OrderStatus: getOrderStatus(order.State),
    }
    return o
}
