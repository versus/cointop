package main

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"fmt"
	"github.com/versus/cointop/source"
)


func main() {

	coinMarketCap, err := source.CoinMarketCapGetData()
	if err != nil {
		panic(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rank", "Coin", "Name", "Price USD", "Price BTC",  "Percent 1h", "Percent 24h", "Percent 7d", "Last update"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(coinMarketCap)
	table.Render()

	source.KrakenGetData()
	fmt.Println("Thats all!")
}