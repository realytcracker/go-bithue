/*
go-bithue by ytcracker
correlate your Philips Hue lights to the bitcoin price on Bitfinex
*/

package main

import (
	"encoding/json"
	//"fmt"
	"github.com/realytcracker/gohue"
	"net/http"
	"os"
	"strconv"
	"time"
)

//Config is loaded from config.json
type Config struct {
	BridgeIP string `json:"BridgeIP"`
	Username string `json:"Username"`
}

//Ticker is a go struct of Bitfinex's json ticker
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
		time.Sleep(15 * time.Second)
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

	lights, _ := bridge.GetAllLights()
	for _, light := range lights {
		light.SetColorHS(color)
	}
}

var bridge hue.Bridge
var ticker Ticker
var config Config

func main() {
	file, _ := os.Open("config.json")
	json.NewDecoder(file).Decode(&config)
	bridge.IPAddress = config.BridgeIP
	bridge.Username = config.Username

	lights, _ := bridge.GetAllLights()
	for _, light := range lights {
		light.SetBrightness(100)
	}

	pollBitstamp()
}
