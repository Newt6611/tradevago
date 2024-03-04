package tri

import "time"

const (
    TimeSignalToTrade time.Duration = time.Second * 5
)

type TradeSignalHandler struct {
    cache map[string]time.Time
}

func NewTradeSignalHandler() *TradeSignalHandler {
    return &TradeSignalHandler {
        cache: map[string]time.Time{},
    }
}

func (t *TradeSignalHandler) StartTradeOrNot(name string) (time.Duration, bool) {
    initTime, ok := t.cache[name]
    if !ok {
        currentTime := time.Now()
        t.cache[name] = currentTime
        initTime = currentTime
    }

    if time.Since(initTime) >= TimeSignalToTrade {
        return time.Since(initTime), true
    }
    return time.Since(initTime), false
}

func (t *TradeSignalHandler) Clear(d time.Duration) {
    ticker := time.NewTicker(d)
    for {
        for k, v := range t.cache {
            if time.Since(v) > TimeSignalToTrade * 2 {
                delete(t.cache, k)
            }
        }
        <-ticker.C
    }
}
