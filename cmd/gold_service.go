package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var currency string = "BRL"

type Gold struct {
	Client *http.Client
	Items  []Item `json:"items"`
}

type Item struct {
	Currency string    `json:"curr"`
	XAUPrice float64   `json:"xauPrice"` // price
	ChgXAU   float64   `json:"chgXau"`   // change
	XAUClose float64   `json:"xauClose"` // previousClose
	Time     time.Time `json:"-"`        // ignored in JSON
}

func (g *Gold) GetPrices() (*Item, error) {
	if g.Client == nil {
		g.Client = &http.Client{}
	}

	url := fmt.Sprintf("https://data-asg.goldprice.org/dbXRates/%s", currency)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("failed to create request:", err)
		return nil, err
	}

	resp, err := g.Client.Do(req)
	if err != nil {
		log.Println("failed to connect with gold price API:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("failed to read response body:", err)
		return nil, err
	}

	var responseData Gold
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		log.Println("failed to unmarshal JSON body:", err)
		return nil, err
	}

	if len(responseData.Items) == 0 {
		log.Println("no items in API response")
		return nil, fmt.Errorf("no items in API response")
	}

	item := responseData.Items[0]
	item.Time = time.Now()

	return &item, nil
}
