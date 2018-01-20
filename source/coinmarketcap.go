package source

import (
	"strconv"
	"fmt"
	"time"
	"encoding/json"

	"github.com/versus/cointop/lib"
)

type CoinMarketCap struct {
	Id	string `json:"id"`
	Name	string `json:"name"`
	Symbol	string `json:"symbol"`
	Rank	string `json:"rank"`
	PriceUsd	string `json:"price_usd"`
	PriceBtc	string `json:"price_btc"`
	Percent1h	string `json:"percent_change_1h"`
	Percent24h	string `json:"percent_change_24h"`
	Percent7d	string `json:"percent_change_7d"`
	LastUpdated	string `json:"last_updated"`
}


// CoinMarketCapGetData func return array of coin

func CoinMarketCapGetData() ([][]string, error) {
	const (
		top10  = 10
		URL = "https://api.coinmarketcap.com/v1/ticker/?limit=10&convert=USD"
	)
	json_data, err := lib.GetJson(URL)
	if err != nil {
		return nil, err
	}
	coins := make([]CoinMarketCap, top10)

	if err := json.Unmarshal(json_data, &coins); err != nil {
		panic(err)
	}

	data := make( [][]string, top10)
	for _, coin := range coins {
		i, err := strconv.ParseInt(coin.LastUpdated, 10, 64)
		if err != nil {
			panic(err)
		}
		item := []string{ coin.Rank,
			coin.Symbol,
			coin.Name,
			coin.PriceUsd,
			coin.PriceBtc,
			coin.Percent1h,
			coin.Percent24h,
			coin.Percent7d,
			fmt.Sprint(time.Unix(i, 0)),
		}
		data = append(data, item)

	}
	return data, err
}