package bito

type bitoHistoryOrder struct {
    ID                string `json:"id"`
    Pair              string `json:"pair"`
    Price             string `json:"price"`
    AvgExecutionPrice string `json:"avgExecutionPrice"`
    Action            string `json:"action"`
    Type              string `json:"type"`
    Timestamp         int64  `json:"timestamp"`
    UpdatedTimestamp  int64  `json:"updatedTimestamp"`
    CreatedTimestamp  int64  `json:"createdTimestamp"`
    Status            int    `json:"status"`
    OriginalAmount    string `json:"originalAmount"`
    RemainingAmount   string `json:"remainingAmount"`
    ExecutedAmount    string `json:"executedAmount"`
    Fee               string `json:"fee"`
    FeeSymbol         string `json:"feeSymbol"`
    BitoFee           string `json:"bitoFee"`
    Total             string `json:"total"`
    Seq               string `json:"seq"`
    TimeInForce       string `json:"timeInForce"`
}

type bitoHistoryEventData struct {
    Event     string                            `json:"event"`
    Timestamp int64                             `json:"timestamp"`
    Datetime  string                            `json:"datetime"`
    Data      map[string][]bitoHistoryOrder     `json:"data"`
}
