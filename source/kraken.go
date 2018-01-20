package source

import (
	"fmt"
	"time"
	"encoding/json"
	"log"
	"github.com/versus/cointop/lib"
)

type TickerResponse struct {
	BCHEUR   PairTickerInfo
	BCHUSD   PairTickerInfo
	BCHXBT   PairTickerInfo
	DASHEUR  PairTickerInfo
	DASHUSD  PairTickerInfo
	DASHXBT  PairTickerInfo
	EOSETH   PairTickerInfo
	EOSEUR   PairTickerInfo
	EOSUSD   PairTickerInfo
	EOSXBT   PairTickerInfo
	GNOETH   PairTickerInfo
	GNOEUR   PairTickerInfo
	GNOUSD   PairTickerInfo
	GNOXBT   PairTickerInfo
	USDTZUSD PairTickerInfo
	XETCXETH PairTickerInfo
	XETCXXBT PairTickerInfo
	XETCZEUR PairTickerInfo
	XETCXUSD PairTickerInfo
	XETHXXBT PairTickerInfo
	XETHZCAD PairTickerInfo
	XETHZEUR PairTickerInfo
	XETHZGBP PairTickerInfo
	XETHZJPY PairTickerInfo
	XETHZUSD PairTickerInfo
	XICNXETH PairTickerInfo
	XICNXXBT PairTickerInfo
	XLTCXXBT PairTickerInfo
	XLTCZEUR PairTickerInfo
	XLTCZUSD PairTickerInfo
	XMLNXETH PairTickerInfo
	XMLNXXBT PairTickerInfo
	XREPXETH PairTickerInfo
	XREPXXBT PairTickerInfo
	XREPZEUR PairTickerInfo
	XREPZUSD PairTickerInfo
	XXBTZCAD PairTickerInfo
	XXBTZEUR PairTickerInfo
	XXBTZGBP PairTickerInfo
	XXBTZJPY PairTickerInfo
	XXBTZUSD PairTickerInfo
	XXDGXXBT PairTickerInfo
	XXLMXXBT PairTickerInfo
	XXLMZEUR PairTickerInfo
	XXLMZUSD PairTickerInfo
	XXMRXXBT PairTickerInfo
	XXMRZEUR PairTickerInfo
	XXMRZUSD PairTickerInfo
	XXRPXXBT PairTickerInfo
	XXRPZCAD PairTickerInfo
	XXRPZEUR PairTickerInfo
	XXRPZJPY PairTickerInfo
	XXRPZUSD PairTickerInfo
	XZECXXBT PairTickerInfo
	XZECZEUR PairTickerInfo
	XZECZUSD PairTickerInfo
}


type KrakenResponse struct {
	Error  []string    `json:"error"`
	Result interface{} `json:"result"`
}


// PairTickerInfo represents ticker information for a pair
type PairTickerInfo struct {
	// Ask array(<price>, <whole lot volume>, <lot volume>)
	Ask []string `json:"a"`
	// Bid array(<price>, <whole lot volume>, <lot volume>)
	Bid []string `json:"b"`
	// Last trade closed array(<price>, <lot volume>)
	Close []string `json:"c"`
	// Volume array(<today>, <last 24 hours>)
	Volume []string `json:"v"`
	// Volume weighted average price array(<today>, <last 24 hours>)
	VolumeAveragePrice []string `json:"p"`
	// Number of trades array(<today>, <last 24 hours>)
	Trades []int `json:"t"`
	// Low array(<today>, <last 24 hours>)
	Low []string `json:"l"`
	// High array(<today>, <last 24 hours>)
	High []string `json:"h"`
	// Today's opening price
	OpeningPrice float32 `json:"o,string"`
}

func KrakenGetData()  {
	const (
		URL = "https://api.kraken.com/0/public/Ticker?pair="
		// https://cex.io/rest-api#ticker
		// https://www.bitstamp.net/api/
	)
	pairs  := []string{  "XXBTZUSD","XETHZUSD", "XLTCZUSD", "XXRPZUSD", "DASHUSD", "XETCZUSD"}
	//var pairResp PairTickerInfo
	var krakenResp KrakenResponse
	//var tikerResp TickerResponse
	for _, pair := range pairs {
		time.Sleep(time.Second * 2)
		json_data, err := lib.GetJson(fmt.Sprint(URL , pair))
		if err != nil {
			fmt.Println("Error ", err.Error())

		}
		if err := json.Unmarshal(json_data, &krakenResp); err != nil {
			log.Println("Error ", err)
		}


		/*
		if err := json.Unmarshal(krakenResp.Result, &pairResp); err != nil {
			log.Println("Error ", err)
		}
		fmt.Println(pairResp)
*/
	}


}

