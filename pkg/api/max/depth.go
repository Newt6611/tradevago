package max

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"strconv"

	"github.com/Newt6611/tradevago/internal"
	"github.com/Newt6611/tradevago/pkg/api"
)

type depthData struct {
	Timestamp         int64      `json:"timestamp"`
	LastUpdateVersion int64      `json:"last_update_version"`
	LastUpdateID      int        `json:"last_update_id"`
	Asks              [][]string `json:"asks"`
	Bids              [][]string `json:"bids"`
}

func (m *Max) GetDepth(ctx context.Context, depthService *api.DepthService) (api.Depth, error) {
    if len(depthService.Pair) <= 0 {
        return api.Depth{}, errors.New("Pair must not be empty")
    }
    query := url.Values{}
    query.Add("market", depthService.Pair)
    query.Add("limit", strconv.Itoa(depthService.Limit))
    query.Add("sort_by_price=", strconv.FormatBool(depthService.SortByPrice))

    res, err := internal.Get(ctx, MAX_API_ENDPOINT + "api/v2/depth", query)
    if err != nil {
        return api.Depth{}, err
    }

    var data depthData
	if err = json.Unmarshal([]byte(res), &data); err != nil {
        return api.Depth{}, err
	}

    return mapDepthData(data), nil
}

func mapDepthData(res depthData) api.Depth {
    var depth api.Depth
    for i := range res.Asks {
        askPrice, _ := strconv.ParseFloat(res.Asks[i][0], 64)
        askAmount, _ := strconv.ParseFloat(res.Asks[i][1], 64)

        bidPrice, _ := strconv.ParseFloat(res.Bids[i][0], 64)
        bidAmount, _ := strconv.ParseFloat(res.Bids[i][1], 64)

        depth.Asks = append(
            depth.Asks,
            api.DepthInfo {
                Price: askPrice,
                Amount: askAmount,
            },
        )

        depth.Bids = append(
            depth.Bids,
            api.DepthInfo {
                Price: bidPrice,
                Amount: bidAmount,
            },
        )
    }
    return depth
}
