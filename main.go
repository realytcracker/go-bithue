/*
go-bithue by ytcracker
correlate your Philips Hue lights to the bitcoin price on Bitstamp
*/

package main

import (
	"encoding/json"
	//"fmt"
	"github.com/realytcracker/gohue"
	"net/http"
	"strconv"
	"time"
)

//BridgeIP is the ip address of your Hue bridge
const BridgeIP = "192.168.1.100"

//Username is a Hue API username for your crib or office or whatever
const Username = "SET ME"

//Ticker is a go struct of bitstamp's json ticker
type Ticker struct {
	Mid       string `json:"mid"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	LastPrice string `json:"last_price"`
	Low       string `json:"low"`
	High      string `json:"high"`
	Volume    string `json:"volume"`
	Timestamp string `json:"timestamp"`
}

func pollBitstamp() {
	for {
		time.Sleep(5 * time.Second)
		getTicker()
	}
}

func getTicker() {
	client := &http.Client{Timeout: time.Second * 5}
	resp, _ := client.Get("https://api.bitfinex.com/v1/pubticker/btcusd")

	json.NewDecoder(resp.Body).Decode(&ticker)

	high, _ := strconv.ParseFloat(ticker.High, 32)
	low, _ := strconv.ParseFloat(ticker.Low, 32)
	lastprice, _ := strconv.ParseFloat(ticker.LastPrice, 32)
	difference := high - low

	floor := lastprice - low
	if floor <= 1 {
		floor = 1
	}

	color := uint16((floor / difference) * 25500)

	/*
		fmt.Printf("lastprice: %.1f\n", lastprice)
		fmt.Printf("low: %.1f\n", low)
		fmt.Printf("high: %.1f\n", high)
		fmt.Printf("difference: %.1f\n", difference)
		fmt.Printf("floor: %.1f\n", floor)
		fmt.Printf("color: %d\n", color)
	*/

	lights, _ := bridge.GetAllLights()
	for _, light := range lights {
		light.SetColorHS(color)
	}
}

var bridge hue.Bridge
var ticker Ticker

func main() {
	bridge.IPAddress = BridgeIP
	bridge.Username = Username

	lights, _ := bridge.GetAllLights()
	for _, light := range lights {
		light.On()
		light.SetBrightness(100)
	}

	pollBitstamp()
}
