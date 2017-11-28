package main

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"time"
	"net/http"
	"bytes"
	"encoding/json"
	"strconv"
	"fmt"
)

const (
	top10  = 10
	URL = "https://api.coinmarketcap.com/v1/ticker/?limit=10&convert=USD"
)

type Coin struct {
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

func getJson(url string) ([]byte, error) {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	return buf.Bytes(), err
}

func getData(url string) ( [][]string, error) {
	json_data, err := getJson(url)
	if err != nil {
		return nil, err
	}
	coins := make([]Coin, top10)

	if err := json.Unmarshal(json_data, &coins); err != nil {
		panic(err)
	}

	data := make( [][]string, top10)
	for _, coin := range coins {
		i, err := strconv.ParseInt(coin.LastUpdated, 10, 64)
		if err != nil {
			panic(err)
		}
		item :=make([]string,9)
		item[0] = coin.Rank
		item[1] = coin.Symbol
		item[2] = coin.Name
		item[3] = coin.PriceUsd
		item[4] = coin.PriceBtc
		item[5] = coin.Percent1h
		item[6] = coin.Percent24h
		item[7] = coin.Percent7d
		item[8] = fmt.Sprint(time.Unix(i, 0))
		data = append(data, item)

	}
	return data, err
}

func main() {

	data, err := getData(URL)
	if err != nil {
		panic(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rank", "Coin", "Name", "Price USD", "Price BTC",  "Percent 1h", "Percent 24h", "Percent 7d", "Last update"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()

}